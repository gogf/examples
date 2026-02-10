---
description: "按照规范梳理示例示例项目源码并创建 README.MD 和 README.ZH.MD 文件。"
agent: "agent"
# model: "Claude Sonnet 4.5"
---

# 前置要求

用户需要提供示例项目的源码路径，如果用户没有提供则需提醒用户提供后再继续后续操作。

# 源码梳理要求

分析用户给定的示例项目的源码后完成以下工作：
- 创建README.MD和README.ZH.MD文件，文档内容格式参考 [文档内容格式指令文件](../instructions/doc-format.instructions.md) 中的要求。
- 如果需要展示目录结构，目录结构的代码高亮语言类型使用text。
- 如果程序需要依赖数据库服务，需要在README文档中需要增加使用docker运行指定服务的代码示例。

