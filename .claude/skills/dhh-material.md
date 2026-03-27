---
name: dhh-material
description: "大航海推送素材工具逻辑"
triggers:
  - "大航海素材工具"
---

# 目的：大航海推送素材页面代码生成

# 技能说明
![页面产品图](./image.png)
![接口逻辑md文档](./dhh-material-api.md)


# 页面字段意义
1. 上传创意页面格局按照：页面产品图 来生成
2. 页面上的字段 和 接口文档中 tool/dhh-material/dhh-forminfo 接口返回的数据得对应
    * 入审类型 - 素材-material， 文案就不要显示了，第一版本不需要文案
    * APP - 对应字段 apps 的数据，显示 name，传递 id
    * 任务类型 - 对应字段 taskTypes， 显示name，传递id
    * 广告类型 - 对应字段 adTypes， 显示 name， 传递 Id
    * 素材卖点 - 对应字段 scenarioTypes，显示name，传递 id
    * 关联商品库or承接页面 - 对应字段 taoBizChannels，显示name， 传递id
    * 热点事件 - 对应字段hotEvents，显示name，传递 Id
    * 底图类型 - 对应字段 baseImageTypes， 显示 name，传递 id
    * 自定义标题 和 自定义文案，是一个用户输入框
3. 选择文件上传的功能，需要调用接口 tool/dhh-material/upload，用户会选定一个文件夹，需要前端从文件夹中循环调用接口上传，
一次接口只能传递10个素材，直到文件夹所有素材上传完毕
 - 每次调用接口，都会返回 metaList 列表，你需要记录每一次的返回列表，在所有素材上传完毕之后，将得到的所有数据以表格展示给用户
4. 展示给用户素材数据之后，用户可以选择提交，提交的饿时候，需要用到 tool/dhh-material/material-create 接口
 - 这个接口中，上传的字段就是第二步选择的字段
  * admissionType 固定为 material
  * appId 传递第二步选择的 APP
  * taskType 传递第二步选择的 任务类型
  * adType 传递第二步选择的 广告类型
  * scenarioType 传递第二步选择的 素材卖点
  * hotEvent 传递第二步选择的 热点事件
  * baseImageType 传递第二步选择的底图类型
  * bizType 传递第二步选择的 关联商品库or承接页面
  * customTitle 和 customCopy 就是输入框的标题和文案
 - materialList 是重点，这里传递的是第三步获取到的 metaList ，详细信息可以看接口文档
 - 注意的是，这个接口materialList一次最多也只能传递 10 个素材信息

# 工具生成
1. 我希望生成一个工具页面，代码打包的时候，windows系统生成一个 exe 文件，mac 生成一个 unix 执行文件
2. 比如windows 系统，用户在点击 exe 文件的时候，能弹出一个工具页面，用来上传素材
3. 工具页面前端展示，按照第二步来，上传数据逻辑可以按照接口逻辑md文档来
4. 我这个功能在另外一个 php 项目上有写，前端代码时 ../../index.vue，后端代码时 ../../DhhMaterialController.php ，可以参考代码
5. 这个工具，我希望用 go 来写