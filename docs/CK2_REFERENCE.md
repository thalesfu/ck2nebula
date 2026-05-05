# CK2 数据库与代码参考文档

> 面向 Neo4j 适配工作的详细技术参考。涵盖现有 Nebula 实现的完整 schema、ORM 机制、查询模式，以及每个概念到 Neo4j 的映射建议。

---

## 目录

1. [模块架构与依赖](#1-模块架构与依赖)
2. [nebulagolang ORM 层](#2-nebulagolang-orm-层)
3. [ck2nebula 数据模型](#3-ck2nebula-数据模型)
   - 3.1 [Vertex Tag 完整 Schema](#31-vertex-tag-完整-schema)
   - 3.2 [Edge 完整 Schema](#32-edge-完整-schema)
4. [VID / EID 格式规范](#4-vid--eid-格式规范)
5. [nGQL 查询模式](#5-ngql-查询模式)
6. [数据来源与解析路径](#6-数据来源与解析路径)
   - 6.1 [数据来源分类](#61-数据来源分类)
   - 6.2 [文件路径常量](#62-文件路径常量)
   - 6.3 [静态数据：从游戏文件解析](#63-静态数据从游戏文件解析)
   - 6.4 [动态数据：从存档文件解析](#64-动态数据从存档文件解析)
   - 6.5 [完整启动流程](#65-完整启动流程)
7. [CK2Commands 脚本生成](#7-ck2commands-脚本生成)
8. [数据库现状统计](#8-数据库现状统计)
9. [Neo4j 映射速查表](#9-neo4j-映射速查表)

---

## 1. 模块架构与依赖

```
paradoxtools          ← 解析 .ck2 存档文件 (SaveFile, Characters, Dynasties...)
     │
     ▼
ck2nebula             ← CK2 实体定义 + 数据转换管道
  tag_*.go            ← 14 个 Vertex 类型
  edge_*.go           ← 41 个 Edge 类型
  people.go           ← 核心生成函数
  space.go            ← 全局 SPACE 初始化
  command/            ← nGQL 查询封装
     │
     ▼
nebulagolang          ← 图 ORM 框架（基于 struct tag 反射）
  nebula.go           ← 连接池管理（max 300）
  space.go            ← Space DDL/DML 操作
  vertex.go           ← 泛型 Vertex CRUD
  edge.go             ← 泛型 Edge CRUD
  vertexcommands.go   ← nGQL 命令生成
  searchvertexcommand.go ← 查询命令生成
     │
     ▼
vesoft-inc/nebula-go  ← Nebula 官方 Go Driver
     │
     ▼
NebulaGraph v3.6      ← 图数据库（Space: "ck2", "ffta", "howiestudy"）


CK2Commands           ← 主程序（占代码量 68%）
  main/main.go        ← flag 解析，调用 ck2nebula + 查询
  scriptbuilder.go    ← 控制台命令文件生成
  family/*/           ← 各家族专属逻辑（婚配、友谊、污染脚本）
```

**go.work 模块列表**

| 模块路径 | 用途 |
|---------|------|
| `CK2Commands` | 主程序 |
| `ck2nebula` | CK2 实体 schema + 数据转换 |
| `fftanebula` | FFTA 实体 schema |
| `golangutils` | 共享工具（文件I/O、图像、日志、SLS） |
| `nebulagolang` | 图 ORM 核心层 |
| `paradoxtools` | Paradox 游戏文件解析 |

---

## 2. nebulagolang ORM 层

### 2.1 Struct Tag 系统

所有实体通过 Go struct tag 描述，反射层在运行时自动生成 nGQL DDL/DML。

#### Vertex Tag

| Go struct tag | 含义 | 示例值 |
|--------------|------|-------|
| `nebulakey:"vid"` | 标识 VID 字段 | `nebulakey:"vid"` |
| `nebulatagname:"xxx"` | 顶点所属的 Tag 名 | `nebulatagname:"people"` |
| `nebulatagcomment:"xxx"` | Tag 的描述（DDL COMMENT） | `nebulatagcomment:"人物"` |
| `nebulaproperty:"xxx"` | Nebula 中的属性名 | `nebulaproperty:"name"` |
| `nebulatype:"xxx"` | Nebula 属性类型（默认 string） | `nebulatype:"Date"`, `nebulatype:"int64"` |
| `nebulaindexes:"xxx"` | 在该字段上建索引 | `nebulaindexes:"name"` |
| `mappingalias:"xxx"` | 从源数据映射时的字段别名 | `mappingalias:"PlayID"` |

#### Edge Tag

| Go struct tag | 含义 |
|--------------|------|
| `nebulakey:"edgefrom"` | 边的起点 VID 字段 |
| `nebulakey:"edgeto"` | 边的终点 VID 字段 |
| `nebulaedgename:"xxx"` | 边类型名 |
| `nebulaedgecomment:"xxx"` | 边描述（DDL COMMENT） |
| `nebulaproperty:"xxx"` | 边属性名 |
| `nebulatype:"xxx"` | 边属性类型 |
| `nebulaindexes:"xxx"` | 边属性索引 |

### 2.2 核心类型

```go
// nebulagolang/space.go
type Space struct {
    Name   string
    Nebula *NebulaDB
}

// nebulagolang/result.go
type Result struct {
    Ok       bool
    Err      error
    Commands []string
    DataSet  *nebulago.ResultSet
}

type ResultT[T any] struct {
    Result *Result
    Data   T
    Ok     bool
    Err    error
}
```

### 2.3 Vertex 泛型操作 API

```go
// vertex.go — 所有函数签名（T 须有 nebulakey:"vid" 和 nebulatagname 标注）

// 写入
InsertVertexes[T](space *Space, vs ...T) *Result
BatchInsertVertexes[T](space *Space, batch int, vs []T) *Result   // 推荐，每批 250
UpdateVertexes[T](space *Space, vs ...T) *Result
UpsertVertexes[T](space *Space, vs ...T) *Result
DeleteVertexes[T](space *Space, vs ...T) *Result
DeleteVertexesByVids(space *Space, vids ...string) *Result
DeleteVertexesWithEdges[T](space *Space, vs ...T) *Result          // 同时删除关联边

// 读取
LoadVertex[T](space *Space, t T) *Result
GetVertexByVid[T](space *Space, vid string) *ResultT[T]
GetAllVertexesByVertexType[T](space *Space) *ResultT[map[string]T]
GetAllVertexesByQuery[T](space *Space, query string) *ResultT[map[string]T]
QueryVertexesByQueryToSlice[T](space *Space, query string) *ResultT[[]T]
GetAllVertexesVIDsByQuery[T](space *Space, query string) *ResultT[map[string]bool]
GetAllVertexesPropertyByQuery[T](space *Space, query, propertyName, displayName string) *ResultT[map[string]bool]

// 工具
GetVID(v interface{}) string
IsVertex[T]() (bool, error)
```

### 2.4 Edge 泛型操作 API

```go
// edge.go — T 须有 nebulakey:"edgefrom"、nebulakey:"edgeto"、nebulaedgename 标注

// 写入
InsertEdges[T](space *Space, es ...T) *Result
BatchInsertEdges[T](space *Space, batch int, es []T) *Result       // 推荐，每批 250
UpdateEdges[T](space *Space, es ...T) *Result
UpsertEdges[T](space *Space, es ...T) *Result
DeleteEdges[T](space *Space, es ...T) *Result
DeleteEdgesByEids(space *Space, eids ...*EID) *Result

// 读取
GetAllEdgesByEdgeType[T](space *Space) *ResultT[map[string]T]
GetAllEdgesByQuery[T](space *Space, query string) *ResultT[map[string]T]
GetEdgeByEid[T](space *Space, eid *EID) *ResultT[T]
GetAllEdgesEIDsByQuery[T](space *Space, query string) *ResultT[map[string]bool]
LoadEdge[T](space *Space, e T) *Result

// 工具
GetEdgeName[T]() string
IsEdge[T]() (bool, error)
```

### 2.5 DDL 命令生成（Space 层）

```go
// space.go 中的 Schema 管理方法
CreateTag(schema *TagSchema) *Result
CreateTagWithIndexes(schema *TagSchema) *Result
RebuildTagWithIndexes(schema *TagSchema) *Result
DropTag(tag string) *Result
AddTagProperty(tag string, prop *TagPropertySchema) *Result
ChangeTagProperty(tag string, prop *TagPropertySchema) *Result

CreateEdge(schema *EdgeSchema) *Result
CreateEdgeWithIndexes(schema *EdgeSchema) *Result
DropEdgeWithIndexes(edge string) *Result

// 生成的 DDL 示例:
// CREATE TAG IF NOT EXISTS people(id int64, name string, ...) COMMENT = "people";
// CREATE TAG INDEX IF NOT EXISTS people_name_index ON people(name(256));
```

### 2.6 nGQL 命令生成（命令层）

```go
// vertexcommands.go
vertexInsertCommand[T](vs ...T) string
// → INSERT VERTEX IF NOT EXISTS people(id,name,...) VALUES "people.123.456":(123,"张三",...)

vertexUpdateCommand[T](v T) string
// → UPDATE VERTEX ON people "people.123.456" SET name="李四" YIELD name AS name

vertexUpsertCommand[T](v T) string
// → UPSERT VERTEX ON people "people.123.456" SET name=$_ YIELD name AS name

// searchvertexcommand.go
LookupTagQueryCommand(t reflect.Type, query string) string
// → LOOKUP ON people [WHERE people.isdead==false] YIELD VERTEX AS v

YieldVertexPropertyNamesCommand(t reflect.Type) string
// → YIELD id($-.v) AS vid, properties($-.v).name AS name, ...

AllVertexesVidsByQueryCommand(t reflect.Type, query string) string
// → LOOKUP ON people WHERE ... YIELD VERTEX AS v | YIELD id($-.v) AS vid

FetchVertexByVidCommand(t reflect.Type, vid string) string
// → FETCH PROP ON people "people.123.456" YIELD VERTEX AS v

CommandPipelineCombine(stmts ...string) string
// 用 | 连接多条语句形成 pipeline
```

---

## 3. ck2nebula 数据模型

### 3.1 Vertex Tag 完整 Schema

#### story（16 个属性）

VID 格式：`story.{story_id}`

| 属性 | 类型 | 说明 |
|------|------|------|
| story_id | int64 | 游戏周目 ID（存档 hash 生成） |
| story_date | date | 游戏内日期 |
| player_id | int64 | 玩家角色 ID |
| player_name | string | 玩家角色全名（含称号） |
| player_age | int64 | 玩家年龄 |
| culture | string | 玩家文化 code |
| culture_name | string | 玩家文化中文名 |
| religion | string | 玩家宗教 code |
| religion_name | string | 玩家宗教中文名 |
| dynasty | string | 玩家家族名 |
| government | string | 政府类型 code |
| government_name | string | 政府类型名 |
| file_path | string | 存档文件路径 |
| file_hash | string | 存档文件 SHA256 |
| file_update_time | datetime | 存档更新时间 |
| version | string | 游戏版本号 |

> **实测数据**：当前仅 1 条 story（story_id=916505602，日期 996-01-01，玩家「女皇 傅 珠珠·不忠者」）

---

#### people（60 个属性）

VID 格式：`people.{character_id}.{story_id}`

**标识与状态**

| 属性 | 类型 | 说明 |
|------|------|------|
| id | int64 | 游戏内角色 ID |
| story_id | int64 | 周目 ID |
| name | string | 名字 |
| nickname | string | 绰号（本地化后） |
| nicknamecode | string | 绰号代码 |
| dna | string | DNA 字符串（外貌） |
| age | int64 | 当前年龄 |
| birthday | date | 生日 |
| deathday | date | 死亡日（存活时为 2000-01-01） |
| isdead | bool | 是否死亡 |
| ishistory | bool | 是否历史人物 |
| isplayer | bool | 是否玩家角色 |
| isbuilt | bool | 是否已写入图库 |
| isbastard | bool | 是否私生子 |
| isfemale | bool | 是否女性 |
| in_hiding | bool | 是否躲藏中 |
| is_pregnant | bool | 是否怀孕 |
| is_consort | bool | 是否为妃妾 |
| married | bool | 是否已婚 |
| consort_count | int64 | 妾的数量 |
| is_under_my_rule | bool | 是否在玩家统治下 |
| gov | string | 政府类型 |

**属性值（基础 + 修正后）**

| 属性 | 类型 | 说明 |
|------|------|------|
| diplomacy | int64 | 外交（基础值） |
| martial | int64 | 武力（基础值） |
| stewardship | int64 | 内政（基础值） |
| intrigue | int64 | 阴谋（基础值） |
| learning | int64 | 学识（基础值） |
| modified_diplomacy | int64 | 外交（含 buff/debuff） |
| modified_martial | int64 | 武力（含 buff/debuff） |
| modified_stewardship | int64 | 内政（含 buff/debuff） |
| modified_intrigue | int64 | 阴谋（含 buff/debuff） |
| modified_learning | int64 | 学识（含 buff/debuff） |
| modified_combat_rating | int64 | 战斗评级（含 buff/debuff） |
| modified_sex_appeal | int64 | 性魅力（含 buff/debuff） |
| health | float | 健康值（基础） |
| modified_health | float | 健康值（含 buff/debuff） |
| fertility | float | 生育力（基础） |
| modified_fertility | float | 生育力（含 buff/debuff） |
| piety | float | 虔诚度 |
| prestige | float | 声望 |
| wealth | float | 财富 |
| score | float | 分数 |

**关系 ID（冗余存储，用于快速访问）**

| 属性 | 类型 | 说明 |
|------|------|------|
| dynasty | int64 | 家族 ID |
| dynasty_name | string | 家族名（冗余） |
| culture | string | 文化 code（冗余） |
| culture_name | string | 文化名（冗余） |
| gfx_culture | string | GFX 文化 code |
| gfx_culture_name | string | GFX 文化名 |
| religion | string | 宗教 code（冗余） |
| religion_name | string | 宗教名（冗余） |
| secret_religion | string | 秘密宗教 code |
| secret_religion_name | string | 秘密宗教名 |
| father | int64 | 父亲 ID |
| mother | int64 | 母亲 ID |
| grandfather | int64 | 父系祖父 ID |
| grandmother | int64 | 父系祖母 ID |
| maternal_grandfather | int64 | 母系祖父 ID |
| maternal_grandmother | int64 | 母系祖母 ID |
| supreme_ruler | int64 | 最高统治者 ID |
| society | int64 | 所属秘密会社 ID |

---

#### dynasty（7 个属性）

VID 格式：`dynasty.{dynasty_id}.{story_id}`

| 属性 | 类型 | 说明 |
|------|------|------|
| id | int64 | 家族 ID |
| story_id | int64 | 周目 ID |
| name | string | 家族名 |
| culture | string | 文化 code |
| religion | string | 宗教 code |
| coat_of_arms_data | string | 纹章数据 |
| coat_of_arms_religion | string | 纹章宗教 |

---

#### title（19 个属性）

VID 格式：`title.{title_code}.{story_id}`

| 属性 | 类型 | 说明 |
|------|------|------|
| id | string | 头衔代码（如 k_china） |
| story_id | int64 | 周目 ID |
| name | string | 头衔名 |
| title | string | 头衔称号 |
| title_female | string | 女性头衔称号 |
| foa | string | 形式称呼 |
| title_level | string | 级别（baron/count/duke/king/emperor） |
| gender | string | 性别限制 |
| active | bool | 是否存在（有人持有） |
| is_custom | bool | 是否自定义 |
| is_dynamic | bool | 是否动态生成 |
| landless | bool | 是否无领地 |
| mercenary | bool | 是否雇佣军 |
| nomad | bool | 是否游牧 |
| rebels | bool | 是否叛乱势力 |
| temporary | bool | 是否临时头衔 |
| major_revolt | bool | 是否重大叛乱 |
| coat_of_arms_data | string | 纹章数据 |
| coat_of_arms_religion | string | 纹章宗教 |

---

#### province（29 个属性）

VID 格式：`province.{province_id}.{story_id}`

| 属性 | 类型 | 说明 |
|------|------|------|
| id | int64 | 省份 ID |
| story_id | int64 | 周目 ID |
| code | string | 省份代码 |
| name | string | 省份名 |
| culture | string | 文化 code |
| culture_name | string | 文化名 |
| religion | string | 宗教 code |
| religion_name | string | 宗教名 |
| max_settlements | int64 | 最大定居点数 |
| primary_settlement | string | 主定居点代码 |
| tech_melee | float | 近战技术 |
| tech_skirmish | float | 轻步兵技术 |
| tech_cavalry | float | 骑兵技术 |
| tech_naval | float | 海军技术 |
| tech_siege_equipment | float | 攻城器械技术 |
| tech_fortifications_construction | float | 防御工事建造技术 |
| tech_castle_construction | float | 城堡建造技术 |
| tech_city_construction | float | 城市建造技术 |
| tech_temple_construction | float | 神庙建造技术 |
| tech_trade_practices | float | 贸易实践技术 |
| tech_infrastructure | float | 基础设施技术 |
| tech_legalism | float | 法制技术 |
| tech_noble_customs | float | 贵族习俗技术 |
| tech_cultural_flexibility | float | 文化灵活性 |
| tech_culture_flex | float | 文化灵活性（冗余） |
| tech_religious_customs | float | 宗教习俗技术 |
| tech_popular_customs | float | 民俗技术 |
| tech_majesty | float | 威严技术 |
| tech_recruitment | float | 征兵技术 |
| tech_infantry | float | 步兵技术 |
| tech_construction | float | 建造技术 |

---

#### baron（7 个属性）

VID 格式：`baron.{baron_code}.{story_id}`

| 属性 | 类型 | 说明 |
|------|------|------|
| code | string | 男爵领代码 |
| story_id | int64 | 周目 ID |
| name | string | 名称 |
| type | string | 类型（castle/city/temple/tribal） |
| leader | int64 | 领主 ID |
| built_date | date | 建成日期 |
| building_date | date | 开始建造日期 |

---

#### culture（13 个属性）

VID 格式：`culture.{culture_code}`（**静态数据，无 story_id**）

| 属性 | 类型 | 说明 |
|------|------|------|
| code | string | 文化代码 |
| name | string | 文化名 |
| allow_looting | bool | 允许劫掠 |
| founder_named_dynasties | bool | 创始人命名家族 |
| dynasty_name_first | bool | 家族名在前 |
| dynasty_title_names | bool | 使用头衔名家族 |
| dukes_called_kings | bool | 公爵称国王 |
| count_titles_hidden | bool | 隐藏伯爵头衔 |
| baron_titles_hidden | bool | 隐藏男爵头衔 |
| caste | bool | 种姓制度 |
| from_dynasty_prefix | string | 来自家族前缀 |
| male_paron_ym | string | 男性父名后缀 |
| female_paron_ym | string | 女性父名后缀 |

> 实测：129 条 culture 数据

---

#### culturegroup（2 个属性）

VID 格式：`culturegroup.{code}`（**静态数据**）

| 属性 | 类型 | 说明 |
|------|------|------|
| code | string | 文化组代码 |
| name | string | 文化组名 |

---

#### religion（2 个属性）

VID 格式：`religion.{religion_code}`（**静态数据**）

| 属性 | 类型 | 说明 |
|------|------|------|
| code | string | 宗教代码 |
| name | string | 宗教名 |

> 实测：52 条 religion 数据

---

#### religiongroup（2 个属性）

VID 格式：`religiongroup.{code}`（**静态数据**）

| 属性 | 类型 | 说明 |
|------|------|------|
| code | string | 宗教组代码 |
| name | string | 宗教组名 |

---

#### trait（130 个属性）

VID 格式：`trait.{trait_code}`（**静态数据**）

核心属性分类（共 130 字段）：

**标识**：id(int64), code(string), name(string), description(string)

**分类布尔标志**：
- 类型：education, personality, virtue, vice, lifestyle, childhood, leader, religious, pilgrimage
- 特殊：immortal, hidden, hidden_from_others, cached, random
- 状态：in_hiding, blinding, incapacitating, is_illness, is_health, is_symptom, is_epidemic
- 遗传：inbred, agnatic, cannot_inherit, cannot_marry, rebel_inherited, succession_gfx
- 其他：customizer, is_close_relative, is_female, is_ruler, is_tribal, is_nomadic, is_theocracy

**属性加成**（int64/float）：diplomacy, martial, stewardship, intrigue, learning, health, fertility, combat_rating, global_levy_size, global_tax_modifier, monthly_character_piety, monthly_character_prestige, monthly_character_wealth, days_of_supply, attrition

**意见值**（int64）：general_opinion, vassal_opinion, same_opinion, dynasty_opinion, liege_opinion, church_opinion, same_religion_opinion, sex_appeal_opinion, spouse_opinion, twin_opinion, 以及各宗教/文化意见

**惩罚值**：diplomacy_penalty, martial_penalty, stewardship_penalty, intrigue_penalty, learning_penalty, health_penalty, fertility_penalty

**继承**：inherit_chance(int64), both_parent_has_trait_inherit_chance(int64), birth(int64)

**AI 参数**：ai_rationality, ai_zeal, ai_greed, ai_honor, ai_ambition

> 实测：444 条 trait 数据

---

#### modifier（105 个属性）

VID 格式：`modifier.{modifier_code}`（**静态数据**）

核心属性：code, name, description, major(bool), icon(int64)

影响类属性（约 100 个，float/int64）：涵盖兵种加成（levy_size, garrison_size, land_morale, 各兵种攻防），税收（local_tax_modifier, global_tax_modifier），建造（build_cost_modifier, build_time_modifier），意见（general_opinion, vassal_opinion 等），技术点（economy_techpoints, military_techpoints, culture_techpoints），人口（population_growth, max_population_mult），贸易（tradevalue, trade_route_wealth），健康（health, fertility, disease_defence），月收益（monthly_character_prestige, monthly_character_piety, monthly_character_wealth）

> 实测：1519 条 modifier 数据

---

#### building（101 个属性）

VID 格式：`building.{building_code}`（**静态数据**）

核心属性：code, name, description, group, group_name, group_description, hasdesc, add_number_to_name(bool), scouting(bool)

建造参数：gold_cost, prestige_cost, piety_cost, build_time, upgrades_from, replaces, convert_to_castle/city/tribal

军事加成：各兵种（light_infantry, heavy_infantry, archers, pikemen, light_cavalry, knights, horse_archers, camel_cavalry, war_elephants）的 offensive/defensive/morale 值，fort_level, garrison_size, levy_size, land_morale, siege_speed/defence

经济加成：tax_income, tradevalue, trade_route_wealth, global_tradevalue, max_tradeposts, levy_reinforce_rate

技术点：building_techpoints, economy_techpoints, military_techpoints

特殊：hospital_level, port(bool), commander_limit, retinuesize, galleys, court_size_modifier

> 实测：465 条 building 数据

---

#### objective（31 个属性）

VID 格式：`objective.{objective_code}`（**静态数据**）

核心属性：code, name, description, text, type, exclusive(bool), can_cancel(bool), cancel_on_leader_death(bool), expectation_of_liege(bool), rel_head_loyalist(bool), ai_capital_kingdom_focus(bool), military_plot(bool), intrigue_plot(bool), murder_plot(bool), vassal_rank_plot(bool), vassal_intrigue_plot(bool)

属性加成：diplomacy, martial, stewardship, intrigue, learning, health, fertility, combat_rating, aggression, plot_power_modifier, warning_level, global_revolt_risk, sex_appeal_opinion, town_opinion, church_opinion

---

### 3.2 Edge 完整 Schema

所有边均有 `story_id(int64)` 属性（静态数据的边无此属性）。

#### Story 汇聚边（6 条，仅 story_id 属性）

| 边类型 | 起点 | 终点 | 说明 |
|--------|------|------|------|
| story_people | story | people | 周目包含的所有人物 |
| story_dynasty | story | dynasty | 周目包含的所有家族 |
| story_title | story | title | 周目包含的所有头衔 |
| story_province | story | province | 周目包含的所有省份 |
| story_baron | story | baron | 周目包含的所有男爵领 |
| story_player | story | people | 周目的玩家角色（单条） |

---

#### People 关系边（17 条）

| 边类型 | 起点 | 终点 | 额外属性 | 说明 |
|--------|------|------|---------|------|
| people_culture | people | culture | — | 人物的文化归属 |
| people_gfx_culture | people | culture | — | 人物的 GFX 文化（外观） |
| people_religion | people | religion | — | 人物的宗教归属 |
| people_secretReligion | people | religion | — | 人物的秘密宗教 |
| people_dynasty | people | dynasty | — | 人物的家族归属 |
| people_trait | people | trait | — | 人物拥有的特性（多条） |
| people_modifier | people | modifier | abate(date), hidden(bool) | 人物身上的 modifier（临时效果） |
| people_claimtitle | people | title | pressed(bool), weak(bool) | 人物的宣称（强主张/弱宣称） |
| people_ambition | people | objective | — | 人物当前的野心 |
| people_focus | people | objective | — | 人物当前的焦点 |
| people_familypeople | people | people | relation(string) | 家族关系（见下方 relation 枚举） |
| people_hostpeople | people | people | — | 人质关系（被扣押方 → 扣押方） |
| people_empirepeople | people | people | — | 帝国关系（属臣 → 皇帝） |
| people_killpeople | people | people | — | 杀害关系（凶手 → 受害者） |
| people_loverpeople | people | people | — | 情人关系 |
| people_guardianpeople | people | people | — | 监护人关系（被监护 → 监护人） |
| people_relatepeople | people | people | code(string), name(string), due(date), detail(string) | 通用关系（盟约、附庸等） |

**people_familypeople.relation 枚举值**

```
father / mother / real_father   ← 亲子关系（real_father 为私生子的真实父亲）
son / daughter                  ← 子女关系
spouse                          ← 配偶
sibling                         ← 兄弟姐妹
```

---

#### Title 关系边（6 条）

| 边类型 | 起点 | 终点 | 额外属性 | 说明 |
|--------|------|------|---------|------|
| title_people | title | people | — | 头衔持有者 |
| title_dynasty | title | dynasty | — | 头衔所属家族 |
| title_basetitle | title | title | — | 基础头衔（男爵领 → 伯爵领 → ...） |
| title_liegetitle | title | title | — | 实际上级头衔 |
| title_dejureliegetitle | title | title | — | 法理上级头衔 |
| title_liegeassimilatingtitle | title | title | de_jure_ass_years(int64) | 正在同化中的上级头衔 |

---

#### Province 关系边（5 条）

| 边类型 | 起点 | 终点 | 额外属性 | 说明 |
|--------|------|------|---------|------|
| province_culture | province | culture | — | 省份的文化 |
| province_religion | province | religion | — | 省份的宗教 |
| province_title | province | title | — | 省份对应的头衔 |
| province_baron | province | baron | — | 省份下辖的男爵领 |
| province_modifier | province | modifier | abate(date), hidden(bool) | 省份上的 modifier |

---

#### Baron 关系边（2 条）

| 边类型 | 起点 | 终点 | 说明 |
|--------|------|------|------|
| baron_title | baron | title | 男爵领对应的头衔 |
| baron_building | baron | building | 男爵领内的建筑 |

---

#### 静态数据分类边（3 条，无 story_id）

| 边类型 | 起点 | 终点 | 说明 |
|--------|------|------|------|
| culturegroup_culture | culturegroup | culture | 文化组包含的文化 |
| religiongroup_religion | religiongroup | religion | 宗教组包含的宗教 |
| trait_oppositetrait | trait | trait | 互斥特性关系 |

---

#### Dynasty 关系边（2 条）

| 边类型 | 起点 | 终点 | 说明 |
|--------|------|------|------|
| dynasty_culture | dynasty | culture | 家族文化 |
| dynasty_religion | dynasty | religion | 家族宗教 |

---

## 4. VID / EID 格式规范

### VID 格式

| Tag | VID 格式 | 示例 | story_id 隔离 |
|-----|---------|------|--------------|
| people | `people.{id}.{story_id}` | `people.2749760.916505602` | ✅ |
| dynasty | `dynasty.{id}.{story_id}` | `dynasty.1000103334.916505602` | ✅ |
| title | `title.{code}.{story_id}` | `title.k_china.916505602` | ✅ |
| province | `province.{id}.{story_id}` | `province.1174.916505602` | ✅ |
| baron | `baron.{code}.{story_id}` | `baron.b_beijing.916505602` | ✅ |
| story | `story.{story_id}` | `story.916505602` | — |
| culture | `culture.{code}` | `culture.han` | ❌ 全局共享 |
| culturegroup | `culturegroup.{code}` | `culturegroup.chinese_group` | ❌ |
| religion | `religion.{code}` | `religion.taoist` | ❌ |
| religiongroup | `religiongroup.{code}` | `religiongroup.chinese_religion` | ❌ |
| trait | `trait.{code}` | `trait.strong` | ❌ |
| modifier | `modifier.{code}` | `modifier.recently_conquered` | ❌ |
| building | `building.{code}` | `building.hillfort_1` | ❌ |
| objective | `objective.{code}` | `objective.obj_become_king` | ❌ |

### EID 格式

Nebula 中边的唯一标识为 `(from_vid, to_vid, edge_type)` 三元组，由 `nebulagolang.EID` 结构封装。代码中通过 `DeleteEdgesByEids` 直接操作。

### 多周目隔离设计

- **动态数据**（people/dynasty/title/province/baron）：VID 中内嵌 `story_id`，天然隔离
- **静态数据**（culture/religion/trait/modifier/building/objective）：全局共享，无 story_id 后缀
- **边**：通过 `story_id` 属性区分不同周目（静态数据间的边无此属性）
- **查询时**：用 `LOOKUP ON people WHERE people.story_id == $storyId` 或直接按 VID 前缀过滤

**注意**：由于 `people.story_id` 等字段在 Nebula 中无索引，`MATCH WHERE n.story_id == x` 会返回空（全表扫描被禁用）。推荐通过 VID 模式（`FETCH PROP ON`）或建索引后使用 `LOOKUP ON`。

---

## 5. nGQL 查询模式

以下是代码库中实际使用的 nGQL 查询模式，Neo4j 适配时需对应翻译为 Cypher。

### 5.1 按特性查找活跃人物（command/peoplesearch.go 核心模式）

```nGQL
-- Step 1: 找到指定特性的 VID
LOOKUP ON trait WHERE trait.code == "strong" YIELD VERTEX AS v

-- Step 2: 提取 VID
| YIELD id($-.v) AS vid

-- Step 3: 反向遍历 people_trait 边，找拥有该特性的活人
| GO FROM $-.vid OVER people_trait REVERSELY
  WHERE properties($$).isdead == false
    AND properties($$).story_id == 916505602
  YIELD $$ AS v,
        properties($$).age AS age,
        properties($$).isfemale AS isfemale

-- Step 4: 排序
| ORDER BY $-.isfemale ASC, $-.age ASC
```

**等价 Cypher**：
```cypher
MATCH (p:people)-[:people_trait]->(t:trait)
WHERE t.code = 'strong'
  AND p.isdead = false
  AND p.story_id = 916505602
RETURN p
ORDER BY p.isfemale ASC, p.age ASC
```

---

### 5.2 遍历家族关系（tag_people.go 中的 nGQL）

```nGQL
-- 从某人出发，找所有家族关系（双向）
GO FROM "people.2749760.916505602" OVER people_familypeople BIDIRECT
  YIELD $$ AS target,
        properties(edge).relation AS relation,
        properties($$).name AS name,
        properties($$).isdead AS isdead
```

**等价 Cypher**：
```cypher
MATCH (p:people {vid: 'people.2749760.916505602'})-[r:people_familypeople]-(other:people)
RETURN other, r.relation AS relation, other.name AS name, other.isdead AS isdead
```

---

### 5.3 查询人物的头衔（tag_people.go）

```nGQL
-- 查询某人持有的所有头衔
GO FROM "people.2749760.916505602" OVER title_people REVERSELY
  YIELD $$ AS title,
        properties($$).name AS title_name,
        properties($$).title_level AS level
```

**等价 Cypher**：
```cypher
MATCH (t:title)-[:title_people]->(p:people {vid: 'people.2749760.916505602'})
RETURN t, t.name AS title_name, t.title_level AS level
```

---

### 5.4 获取全量人物属性（FETCH 模式）

```nGQL
FETCH PROP ON people "people.2749760.916505602" YIELD VERTEX AS v
| YIELD id($-.v) AS vid,
        properties($-.v).name AS name,
        properties($-.v).age AS age,
        properties($-.v).isdead AS isdead,
        properties($-.v).diplomacy AS diplomacy
        -- ... 其余所有属性
```

**等价 Cypher**：
```cypher
MATCH (p:people {vid: 'people.2749760.916505602'})
RETURN p
```

---

### 5.5 dynasty 查询（tag_dynasty.go 中的 nGQL）

```nGQL
-- 查找某家族的所有成员（带年龄性别）
LOOKUP ON people
  WHERE people.dynasty == 1000103334
    AND people.story_id == 916505602
    AND people.isdead == false
  YIELD VERTEX AS v
| YIELD id($-.v) AS vid,
        properties($-.v).name AS name,
        properties($-.v).age AS age,
        properties($-.v).isfemale AS isfemale
| ORDER BY $-.isfemale ASC, $-.age ASC
```

**等价 Cypher**：
```cypher
MATCH (p:people)-[:people_dynasty]->(d:dynasty)
WHERE d.id = 1000103334 AND p.story_id = 916505602 AND p.isdead = false
RETURN p
ORDER BY p.isfemale ASC, p.age ASC
```

---

### 5.6 多跳遍历（1-2 跳）

```nGQL
-- 从玩家出发，遍历家族关系 1-2 跳（命令生成常用模式）
GO 1 TO 2 STEPS FROM "people.2749760.916505602" OVER people_familypeople
  WHERE properties($$).isdead == false
  YIELD DISTINCT $$ AS v
```

**等价 Cypher**：
```cypher
MATCH (start:people {vid: 'people.2749760.916505602'})-[:people_familypeople*1..2]-(p:people)
WHERE p.isdead = false
RETURN DISTINCT p
```

---

### 5.7 properties($$) 模式说明

nGQL 中 `$$` 表示终点顶点的属性，`$^` 表示起点，`$-.fieldName` 表示管道前序结果的字段。

| nGQL 表达式 | 含义 | Cypher 等价 |
|------------|------|------------|
| `properties($$).name` | 终点的 name 属性 | `endNode.name` |
| `properties($^).name` | 起点的 name 属性 | `startNode.name` |
| `properties(edge).relation` | 边的 relation 属性 | `r.relation` |
| `$-.vid` | 管道输入的 vid 列 | `p.vid`（Cypher 无需此语法） |
| `id($-.v)` | 提取顶点 VID | `p.vid`（自定义属性） |

---

## 6. 数据来源与解析路径

### 6.1 数据来源分类

图数据库中的实体分为两类，来源完全不同：

| 类别 | 实体 | 数据源 | 更新时机 |
|------|------|--------|---------|
| **静态数据** | culture, culturegroup, religion, religiongroup, trait, modifier, building, objective | CK2 游戏安装目录 `/common/` 下的 `.txt` 文件 | 手动运行 `-bs` flag |
| **动态数据** | story, people, dynasty, title, province, baron | CK2 存档文件 `.ck2` | 存档变化时自动或手动触发 |

### 6.2 文件路径常量

```go
// CK2Commands/main/main.go
const ck2Folder  = "/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II"
const saveFolder = "/Users/thalesfu/Documents/Paradox Interactive/Crusader Kings II/save games"
const saveFolder2 = "/Users/thalesfu/Windows/steam/userdata/94993760/203770/remote/save games"
```

游戏目录下被读取的子目录结构：

```
{ck2Folder}/
└── common/
    ├── cultures/         *.txt  → Culture + CultureGroup
    ├── religions/        *.txt  → Religion + ReligionGroup
    ├── traits/           *.txt  → Trait + Trait_OppositeTrait
    ├── buildings/        *.txt  → Building
    ├── event_modifiers/  *.txt  → Modifier（主要来源）
    ├── trade_routes/     *.txt  → Modifier（贸易路线相关）
    └── objectives/       *.txt  → Objective
```

存档文件目录：

```
{saveFolder}/           ← 本地存档
{saveFolder2}/          ← Steam 云同步存档（当前实际使用）
    autosave.ck2        ← 当前数据库对应的存档
```

**当前数据库所用存档**

| 字段 | 值 |
|------|---|
| 文件路径 | `/Users/thalesfu/Windows/steam/userdata/94993760/203770/remote/save games/autosave.ck2` |
| 更新时间 | 2024-11-01 11:14:31 |
| story_id | 916505602 |
| 文件 SHA256 | `7612bc4866e18a12e0c84df99cc7b3250c2a4741bc8a382b27b1ca9f5581d2d5...` |

---

### 6.3 静态数据：从游戏文件解析

静态数据通过 `-bs` 标志触发，调用链如下：

```
CK2Commands -bs
  ├─ culture.BuildCulture(ck2Folder)
  │   └─ ck2nebula.GenerateCultureData(path)
  │       └─ paradoxtools/CK2/culture.LoadAllCultures(path)
  │           └─ 读取 {path}/common/cultures/*.txt
  │               pserialize.UnmarshalP[CultureGroup](content)
  │           → []*CultureGroup, []*Culture, []*CultureGroup_Culture
  │
  ├─ religion.BuildReligion(ck2Folder)
  │   └─ ck2nebula.GenerateReligionData(path)
  │       └─ paradoxtools/CK2/religion.LoadAllReligions(path)
  │           └─ 读取 {path}/common/religions/*.txt
  │           → []*ReligionGroup, []*Religion, []*ReligionGroup_Religion
  │
  ├─ ck2nebula.BuildModifiers(ck2Folder)
  │   └─ ck2nebula.GenerateModifiersData(path)
  │       ├─ paradoxtools/CK2/traderoute.LoadAllTradeRoutes(path)
  │       │   └─ 读取 {path}/common/trade_routes/*.txt
  │       └─ paradoxtools/CK2/modifier.LoadAllModifier(path)
  │           └─ 读取 {path}/common/event_modifiers/*.txt
  │           → []*Modifier
  │
  ├─ ck2nebula.BuildObjectives(ck2Folder)
  │   └─ ck2nebula.GenerateObjectivesData(path)
  │       └─ paradoxtools/CK2/objectives.LoadAllObjective(path)
  │           └─ 读取 {path}/common/objectives/*.txt
  │           → []*Objective
  │
  ├─ ck2nebula.BuildBuildings(ck2Folder)
  │   └─ ck2nebula.GenerateBuildingsData(path)
  │       └─ paradoxtools/CK2/building.LoadAllBuildings(path)
  │           └─ 读取 {path}/common/buildings/*.txt
  │           → []*Building
  │
  └─ ck2nebula.BuildTraits(ck2Folder)
      └─ ck2nebula.GenerateTraitsData(path)
          └─ paradoxtools/CK2/trait.LoadAllTraits(path)
              └─ 读取 {path}/common/traits/*.txt
              → []*Trait, []*Trait_OppositeTrait
```

#### 各静态实体的解析细节

**Culture / CultureGroup**

| 项目 | 内容 |
|------|------|
| 解析器 | `paradoxtools/CK2/culture/culture.go` → `LoadAllCultures()` |
| 输入文件 | `common/cultures/*.txt`（Paradox 嵌套文本格式） |
| 文件结构 | 顶层为 CultureGroup block，内嵌多个 Culture block |
| 本地化 | 调用 `localisation.LoadAllTranslations()` 填充 `name` 字段 |
| 输出 | `[]*CultureGroup`, `[]*Culture`, `[]*CultureGroup_Culture` |

**Religion / ReligionGroup**

| 项目 | 内容 |
|------|------|
| 解析器 | `paradoxtools/CK2/religion/religion.go` → `LoadAllReligions()` |
| 输入文件 | `common/religions/*.txt` |
| 文件结构 | 顶层为 ReligionGroup block，内嵌多个 Religion block |
| 输出 | `[]*ReligionGroup`, `[]*Religion`, `[]*ReligionGroup_Religion` |

**Trait**

| 项目 | 内容 |
|------|------|
| 解析器 | `paradoxtools/CK2/trait/trait.go` → `LoadAllTraits()` |
| 输入文件 | `common/traits/*.txt` |
| 特殊处理 | traits 文件按顺序编号，读取顺序即为 `id` 字段；`opposites` 字段展开为 `Trait_OppositeTrait` 边 |
| 输出 | `[]*Trait`（130 个属性/trait），`[]*Trait_OppositeTrait` |

**Modifier**

| 项目 | 内容 |
|------|------|
| 解析器 | `paradoxtools/CK2/modifier/modifier.go` → `LoadAllModifier()` |
| 输入文件 | `common/event_modifiers/*.txt`（主要），`common/trade_routes/*.txt`（补充） |
| 说明 | trade_routes 文件解析后也转为 Modifier 实体存入同一 Tag |
| 输出 | `[]*Modifier`（105 个属性/modifier） |

**Building**

| 项目 | 内容 |
|------|------|
| 解析器 | `paradoxtools/CK2/building/building.go` → `LoadAllBuildings()` |
| 输入文件 | `common/buildings/*.txt` |
| 文件结构 | 顶层为 BuildingGroup，内嵌各等级建筑；upgrades_from/replaces 字段记录升级链 |
| 输出 | `[]*Building`（101 个属性/building） |

**Objective**

| 项目 | 内容 |
|------|------|
| 解析器 | `paradoxtools/CK2/objectives/objective.go` → `LoadAllObjective()` |
| 输入文件 | `common/objectives/*.txt` |
| 说明 | 涵盖野心（ambition）和专注（focus）两类 objective |
| 输出 | `[]*Objective` |

#### Paradox 文本格式解析器

所有静态文件均为 Paradox 专有文本格式，由 `paradoxtools/utils/pserialize` 解析：

```go
// 核心函数
pserialize.UnmarshalP[T](content string) (T, error)

// 支持的 struct tag
`paradox_field:"field_name"`         // 映射字段名
`paradox_type:"map|list|entity"`     // 值类型
`paradox_text:"yes"`                 // 取文本值
`paradox_map_key_pattern:"pattern"`  // map key 匹配模式
```

---

### 6.4 动态数据：从存档文件解析

动态数据在以下情况触发：手动 `-f` flag、`-w` 监听模式检测到文件变化、或首次启动。

#### 存档文件处理

```go
// paradoxtools/CK2/save/savefile.go
save.LoadSave(ck2Folder string, savePath string) (*SaveFile, bool, error)

// 流程：
// 1. 读取文件，检测是否为 zip 压缩格式
// 2. 若压缩：golangutils.Unzip() 解压至 /tmp/ck2/unzipsavefile/{存档名}/
// 3. pserialize.UnmarshalP[SaveFile](content) 解析文本内容
// 4. 内部调用 processTitles()、processProvinces()、processDynasties()、processCharacters()
```

**SaveFile 核心结构**

```go
type SaveFile struct {
    Version    string
    Date       pserialize.Year            // 当前游戏日期
    Player     *Player                    // 玩家信息
    PlayerName string
    Characters map[int]*Character         // 所有角色（含历史/死亡）
    Dynasties  map[int]*Dynasty           // 所有家族
    Titles     map[string]*Title          // 所有头衔（以 title code 为 key）
    Provinces  map[int]*Province          // 所有省份（以 province ID 为 key）
    Relations  map[string]map[string]*PersonRelation  // 外交关系
    Religions  map[string]*Religion       // 存档中的宗教状态
    ActiveAmbitions []*ActiveEvent        // 当前活跃野心
    ActiveFocuses   []*ActiveEvent        // 当前活跃专注
}
```

#### People / 相关边（从 SaveFile.Characters 生成）

```go
// ck2nebula/people.go
func GeneratePeople(
    file            *save.SaveFile,
    cultureMap      map[string]*Culture,         // 静态：文化 code → Culture 实体
    religionMap     map[string]*Religion,        // 静态：宗教 code → Religion 实体
    historyPeople   map[int]map[int]*People,     // 历史人物缓存
    historyDynasty  map[int]*Dynasty,            // 历史家族缓存
    supremeRuler    map[int]int,                 // 人物 → 最高统治者映射
    rulerChain      map[int]map[int]bool,        // 统治链
) (
    []*People,
    []*People_Culture, []*People_GFXCulture,
    []*People_Religion, []*People_SecretReligion,
    []*People_Trait, []*People_Modifier,
    []*People_ClaimTitle, []*People_Dynasty,
    []*People_FamilyPeople, []*People_HostPeople,
    []*People_EmpirePeople, []*People_KillPeople,
    []*People_LoverPeople, []*People_GuardianPeople,
    []*People_Ambition, []*People_Focus,
)
```

**save.Character 核心字段**（存档中每个角色的原始数据）

```go
type Character struct {
    ID             int                  // 游戏内唯一 ID
    Name           string               // 名字（游戏内文本，可能为本地化 code）
    Birthday       pserialize.Year      // 出生日期
    DeathDay       pserialize.Year      // 死亡日期
    IsFemale       pserialize.PBool
    Culture        string               // 文化 code
    GFXCulture     string               // 外观文化 code
    Religion       string               // 宗教 code
    SecretReligion string               // 秘密宗教 code
    Dynasty        int                  // 家族 ID
    Father         int                  // 父亲 ID
    Mother         int                  // 母亲 ID
    RealFather     int                  // 真实父亲 ID（私生子）
    Spouse         []int                // 配偶 ID 列表
    Consort        []int                // 妾 ID 列表
    ConsortOf      int                  // 正妻 ID（是妾时）
    Lover          []int                // 情人 ID 列表
    Host           int                  // 扣押者 ID（人质时）
    Guardian       int                  // 监护人 ID
    Emp            int                  // 帝国主君 ID
    Killer         int                  // 杀害者 ID
    Traits         []int                // 特性 ID 列表（与 trait 静态数据的 id 对应）
    Modifier       []*CharModifier      // modifier 列表（含 abate 日期、hidden 标志）
    Claim          []*Claim             // 宣称头衔列表（含 pressed、weak 标志）
    Titles         []string             // 持有头衔 code 列表
    Diplomacy      int                  // 外交基础值
    Martial        int                  // 武力基础值
    Stewardship    int                  // 内政基础值
    Intrigue       int                  // 阴谋基础值
    Learning       int                  // 学识基础值
    Health         float32              // 健康
    Fertility      float32              // 生育力
    Prestige       float32              // 声望
    Piety          float32              // 虔诚
    Wealth         float32              // 财富
    Score          float32              // 分数
    Society        int                  // 秘密会社 ID
    Gov            string               // 政府类型
    // ... 更多字段
}
```

**GeneratePeople 关键处理逻辑**

1. 遍历 `file.Characters`，为每个角色创建 `People` vertex
2. 从数据库查询 `GetAllModifier(SPACE)` 得到 modifier map，按 code 查找计算 `modified_*` 属性
3. 从数据库查询 `GetAllTraitsMap(SPACE)` 得到 trait map，按 ID 查找计算 `modified_*` 属性
4. 从 `file.ActiveAmbitions` / `file.ActiveFocuses` 查找当前野心/专注，生成对应边
5. 按 `Character.Spouse/Consort/Lover/Guardian/Host/Emp/Killer` 生成各类 People→People 边
6. 通过 `cultureMap`/`religionMap` 填充 `culture_name`/`religion_name` 冗余字段
7. 计算 `grandfather/grandmother/maternal_grandfather/maternal_grandmother`（从父母 ID 递推）

#### Dynasty（从 SaveFile.Dynasties 生成）

```go
// ck2nebula/dynasties.go → GenerateDynasties(dynasties map[int]*save.Dynasty)
// save.Dynasty 字段: ID, Name, Culture, Religion, CoatOfArmsData, CoatOfArmsReligion
// 输出: []*Dynasty, []*Dynasty_Culture, []*Dynasty_Religion
```

#### Title（从 SaveFile.Titles 生成）

```go
// ck2nebula/titles.go → GenerateTitles(titles map[string]*save.Title)
// save.Title 字段: ID(code), Name, TitleFemale, TitleLevel, Gender, Active,
//                  IsCustom, IsDynamic, Landless, Nomad, Rebels, Temporary,
//                  MajorRevolt, Mercenary, CoatOfArmsData, Liege, BaseTitles,
//                  DejureLiegeTitles, AssimilatingLiegeTitles, Holder, Dynasty
// 输出: []*Title, []*Title_LiegeTitle, []*Title_BaseTitle,
//       []*Title_DejureLiegeTitle, []*Title_LiegeAssimilatingTitle,
//       []*Title_Dynasty, []*Title_People
```

#### Province + Baron（从 SaveFile.Provinces 生成）

```go
// ck2nebula/provinces.go → GenerateProvinces(provinces map[int]*save.Province, ...)
// save.Province 字段: ID, Name, Culture, Religion, MaxSettlements, PrimarySettlement,
//                     TechLevels(16种), Modifiers, BaronialHolding(含建筑列表)
// 输出: []*Province, []*Province_Culture, []*Province_Religion,
//       []*Province_Modifier, []*Province_Title,
//       []*Baron, []*Province_Baron,
//       []*Baron_Title, []*Baron_Building
```

Baron 嵌套在 Province 内部：`Province.BaronialHolding` 是 map，每个 key 为男爵领 code，value 含 type、leader、built_date、building_date、建筑列表。

#### Story（汇聚所有动态实体）

```go
// ck2nebula/story.go → GenerateStoryDetails(ck2Folder, savePath, cultureMap, religionMap, ...)
// 流程：
// 1. save.LoadSave() 解析存档
// 2. 调用 GenerateTitles / GenerateProvinces / GenerateDynasties / GeneratePeople
// 3. 创建 Story vertex（从 SaveFile.Player + 文件元数据）
// 4. 生成 story_people / story_dynasty / story_title / story_province / story_baron / story_player 边
```

---

### 6.5 完整启动流程

```
CK2Commands 启动
  │
  ├─── -bs 分支（静态数据，手动运行一次）
  │     │
  │     ├─ culture.LoadAllCultures({ck2Folder}/common/cultures/*.txt)
  │     │   → BatchInsertVertexes(CultureGroups) + BatchInsertEdges(CultureGroup_Cultures)
  │     │   → BatchInsertVertexes(Cultures)
  │     │
  │     ├─ religion.LoadAllReligions({ck2Folder}/common/religions/*.txt)
  │     │   → BatchInsertVertexes(ReligionGroups + Religions + edges)
  │     │
  │     ├─ modifier.LoadAllModifier({ck2Folder}/common/event_modifiers/*.txt)
  │     │   + traderoute.LoadAllTradeRoutes({ck2Folder}/common/trade_routes/*.txt)
  │     │   → BatchInsertVertexes(Modifiers)
  │     │
  │     ├─ objectives.LoadAllObjective({ck2Folder}/common/objectives/*.txt)
  │     │   → BatchInsertVertexes(Objectives)
  │     │
  │     ├─ building.LoadAllBuildings({ck2Folder}/common/buildings/*.txt)
  │     │   → BatchInsertVertexes(Buildings)
  │     │
  │     └─ trait.LoadAllTraits({ck2Folder}/common/traits/*.txt)
  │         → BatchInsertVertexes(Traits) + BatchInsertEdges(Trait_OppositeTraits)
  │
  └─── 动态数据分支（-f 强制 或 -w 监听模式）
        │
        ├─ 扫描 saveFolder + saveFolder2，找最新 .ck2 文件
        │
        ├─ save.LoadSave(ck2Folder, {autosave.ck2 路径})
        │   ├─ 解压（若 zip）→ /tmp/ck2/unzipsavefile/
        │   └─ pserialize.UnmarshalP[SaveFile](content)
        │       SaveFile{ Characters(14万+), Dynasties, Titles, Provinces, Relations }
        │
        ├─ GenerateTitles(file.Titles)
        │   → []*Title + title关系边 × 6种
        │
        ├─ GenerateProvinces(file.Provinces, cultureMap, religionMap)
        │   → []*Province + []*Baron + province/baron关系边 × 7种
        │
        ├─ GenerateDynasties(file.Dynasties)
        │   → []*Dynasty + dynasty关系边 × 2种
        │
        ├─ GeneratePeople(file, cultureMap, religionMap, ...)
        │   ├─ GetAllModifier(SPACE) ← 查询 Nebula，得 modifier map
        │   ├─ GetAllTraitsMap(SPACE) ← 查询 Nebula，得 trait map
        │   └─ 遍历 file.Characters（14万+）
        │       → []*People + people关系边 × 17种
        │
        ├─ GenerateStory(file) → []*Story + story关系边 × 6种
        │
        ├─ CompareAndUpdateNebulaEntityBySliceAndQuery()
        │   ├─ 对比旧数据：Added / Updated / Deleted / Kept
        │   ├─ BatchInsertVertexes(batch=250) × 所有顶点类型
        │   ├─ BatchInsertEdges(batch=250) × 所有边类型
        │   └─ 并发 goroutine（最多 3 个）执行各类型更新
        │
        └─ 查询 Nebula → 生成 CK2 控制台命令脚本文件
```

---

## 7. CK2Commands 脚本生成

### 7.1 启动模式（main/main.go flags）

| Flag | 说明 |
|------|------|
| `-bs` | 构建静态数据（culture/religion/modifier/objective/building/trait） |
| `-f` | 强制重新加载存档数据 |
| `-w` | 监视模式，fsnotify 监听存档文件变化自动触发 |
| `-t` | 头衔模式：生成头衔相关命令脚本 |
| `-c` | 文化模式：生成文化相关命令脚本 |
| `-r` | 宗教模式：生成宗教相关命令脚本 |
| `-n` | 名字模式：生成改名命令脚本 |
| `-wr` | 宗教清洗模式：生成批量改宗脚本 |
| `-m` | 婚配模式：生成婚配命令脚本 |

### 7.2 脚本生成框架（scriptbuilder.go）

```go
type ScriptGenerator interface {
    Generate() []string  // 返回 CK2 控制台命令行列表
}

func BuildScript(fileName string, generators ...ScriptGenerator) {
    // 写入文件，末尾追加 sound_effect = secret_cults_gain_sympathy_01
}
```

### 7.3 核心家族常量

```go
var CoreFamily = map[int]string{
    1000103393: "lou",   // 楼氏
    1000103382: "yuan",  // 袁氏
    1000103379: "chu",   // 褚氏
    // ... 共 9 个家族
}
```

### 7.4 家族模块结构（family/{name}/）

```
family/lou/
  loupeople.go      ← 人物 ID 常量：const M楼彦玮2611890 = 2611890
  loupollinate.go   ← 通婚脚本：Pollinate() → people.Pollinate(couples, "lou")
  loufriends.go     ← 友谊列表：GetLouFriends() []int
```

---

## 8. 数据库现状统计

**运行环境**：NebulaGraph v3.6，Docker Compose，3 metad + 3 graphd + 3 storaged（本地集群）

**当前存档**：`autosave.ck2`（Steam 云同步，2024-11-01 11:14:31，story_id=916505602，游戏日期 996-01-01）

| 实体类型 | 数量 | 数据来源 |
|---------|------|---------|
| story | 1 | 存档文件 |
| people | 140,390 | 存档文件 |
| dynasty | — | 存档文件 |
| title | — | 存档文件 |
| province | — | 存档文件 |
| baron | — | 存档文件 |
| culture | 129 | 游戏文件 `common/cultures/` |
| religion | 52 | 游戏文件 `common/religions/` |
| trait | 444 | 游戏文件 `common/traits/` |
| building | 465 | 游戏文件 `common/buildings/` |
| modifier | 1,519 | 游戏文件 `common/event_modifiers/` + `common/trade_routes/` |

**玩家角色 people.2749760.916505602（珠珠）**

| 属性 | 值 |
|------|---|
| 性别/年龄 | 女，36 岁（生于 0959-10-22） |
| 文化/宗教 | 汉 / 道教 |
| 家族 | 傅氏（dynasty_id=1000103334） |
| 外交/阴谋/学识/内政（基础） | 48 / 48 / 48 / 48 |
| 外交/阴谋/学识/内政（修正后） | 58 / 54 / 63 / 59 |
| 武力（基础/修正后） | 29 / 46 |
| 生命值（基础/修正后） | 9.59 / 11.09 |
| 声望/虔诚/财富 | 5165 / 2901 / 300万+ |
| 当前状态 | 怀孕中，在位女皇，是最高统治者 |

---

## 9. Neo4j 映射速查表

### 9.1 核心概念映射

| Nebula 概念 | Neo4j 对应 | 实现备注 |
|------------|-----------|---------|
| Space `"ck2"` | Database `ck2` | 需 Neo4j 4.x 企业版多数据库；社区版用单库+Label前缀 |
| Tag `people` | Label `:People` | Label 首字母大写是 Neo4j 惯例 |
| VID `"people.123.456"` | 节点属性 `vid` + 唯一约束 | Neo4j 内部 ID 不稳定，必须自存 `vid` |
| Edge `people_culture` | Relationship `[:PEOPLE_CULTURE]` | Neo4j 关系类型用全大写 |
| EID `(from, to, type)` | (起点vid, 终点vid, 类型) | 用 MERGE 三元组定位 |
| `nebulaproperty:"name"` | 属性 `name` | tag 可直接复用，无需修改 |
| `nebulaindexes:"x"` | `CREATE INDEX FOR (n:Tag) ON (n.x)` | 索引命令不同，语义相同 |
| `nebulakey:"vid"` | 用于 `MERGE (n {vid: $vid})` | 复用同一个 tag 字段 |
| `LOOKUP ON trait WHERE ...` | `MATCH (t:Trait) WHERE ...` | Cypher 无需预先建索引才能过滤 |
| `GO FROM vid OVER edge` | `MATCH (a {vid:$v})-[:EDGE]->(b)` | 直接 MATCH，更简洁 |
| `BIDIRECT` | `-[:EDGE]-`（无方向箭头） | Cypher 用无向关系 |
| `properties($$).x` | `endNode.x` | Cypher 直接引用节点变量 |
| `$-.field` | 无需，Cypher 变量在整个查询中共享 | Pipeline 语法差异最大 |
| `ORDER BY $-.col ASC` | `ORDER BY col ASC` | 语法简化 |

### 9.2 VID 在 Neo4j 中的处理

```cypher
-- 创建唯一约束（替代 Nebula 的 VID 唯一性保证）
CREATE CONSTRAINT people_vid IF NOT EXISTS
  FOR (n:People) REQUIRE n.vid IS UNIQUE;

-- UPSERT 写入（等价 Nebula 的 UPSERT VERTEX）
MERGE (n:People {vid: $vid})
SET n += $props

-- 按 VID 查询（等价 FETCH PROP ON people "vid"）
MATCH (n:People {vid: $vid})
RETURN n
```

### 9.3 story_id 多周目隔离

```cypher
-- Nebula: LOOKUP ON people WHERE people.story_id == 916505602
-- Cypher 等价（可建索引加速）:
MATCH (p:People)
WHERE p.story_id = 916505602
RETURN p

-- 或通过边过滤（更图风格）:
MATCH (s:Story {story_id: 916505602})-[:STORY_PEOPLE]->(p:People)
RETURN p
```

### 9.4 需要重写的 nGQL 查询（按复杂度）

| 来源文件 | 查询类型 | 重写难度 |
|---------|---------|---------|
| `command/peoplesearch.go` | LOOKUP + GO REVERSELY + ORDER BY | 中（3步 pipeline → 1条 Cypher） |
| `tag_people.go:160-236` | GO BIDIRECT + properties($$) | 低（MATCH 无向关系） |
| `tag_dynasty.go:78-84` | LOOKUP + YIELD | 低（MATCH WHERE） |
| `command/peoplesearch.go:15-19` | GO REVERSELY + 多条件 WHERE | 中 |

### 9.5 struct tag 复用策略

`neo4jgolang` 新增以下 tag（复用 `nebulaproperty`、`nebulakey`、`nebulaindexes`，新增 Neo4j 专属）：

| 新增 tag | 用途 | 示例 |
|---------|------|------|
| `neo4jlabel:"People"` | 节点 Label（替代 `nebulatagname`） | `neo4jlabel:"People"` |
| `neo4jrel:"PEOPLE_CULTURE"` | 关系类型（替代 `nebulaedgename`） | `neo4jrel:"PEOPLE_CULTURE"` |

`nebulaproperty`、`nebulakey`（vid/edgefrom/edgeto）、`nebulaindexes` 直接在 neo4jgolang 中读取复用，无需重复定义。

---

*文档基于 NebulaGraph v3.6 实例数据（story_id=916505602，存档 autosave.ck2，2024-11-01）和代码库快照（2026-05-04）生成。*
