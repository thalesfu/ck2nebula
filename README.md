# ck2nebula

CK2（十字军之王 II）实体 schema 定义，以及向 Nebula 图数据库的写入路径。

## 模块职责

- `tag_*.go`：节点类型定义（People、Dynasty、Title、Province、Baron、Culture、Religion、Trait 等）
- `edge_*.go`：关系类型定义（People_Trait、People_Culture、Title_People 等，共 17 种关系）
- `people.go`：主数据管道 `GeneratePeople()`，从存档计算 modified_* 属性
- `story.go`：存档写入入口 `BuildStory()` / `LoadAndUpdateStory()`
- `space.go`：初始化 Nebula `ck2` space

## 关键约定

- VID 格式：`people.{id}.{storyId}`，`culture.{code}`，`trait.{code}` 等
- `StoryID` 属性在动态数据上隔离多周目
- `nebulakey:"vid"` 为节点唯一键，`nebulakey:"edgefrom"/"edgeto"` 为边端点
- `modified_*` 字段 = 基础属性 + 所有 trait/modifier bonus 之和

## 配置

模块目录下放 `nebula-account.yaml`（已加入 .gitignore）。

完整实体 / 关系 schema 参考见 [docs/CK2_REFERENCE.md](docs/CK2_REFERENCE.md)。
