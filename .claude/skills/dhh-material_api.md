---
name: dhh-material
description: "大航海推送素材接口"
triggers:
  - "大航海素材"
---

## 目的
此功能的目的主要是为了通过接口给大航海平台推送素材


# 接口文件创建
1. 在 ./modules/tool/controllers/ 这个文件夹下创建：DhhMaterialController.php 文件，用来写给前端调用的接口
2. 我还会需要用到 ./modules/tool/models/ 这个文件夹下的，AdToolGyxToken.php 文件，这个是数据表文件，关于这个表的sql逻辑都写在这个文件里
- 这个项目其他的文件应该都不太需要


# 大航海推送素材逻辑前置接口，此接口需要在 DhhMaterialController.php 中添加
- 我给你提供一个获取前置数据的接口，你需要把这个接口逻辑写在我的文件里，源接口curl为：
```
curl 'https://dhh.taobao.com/polystar/api/creative/material/forminfo?_csrf=1b912238-b177-4f14-a0e0-b8abf9aefc02' \
  -H 'accept: application/json, text/plain, */*' \
  -H 'accept-language: zh-CN,zh;q=0.9' \
  -b 't=aad3d47a3693faf2be2983e66f9366c7; wk_cookie2=167f0c4d672ac968e0a93b99b0d0321b; cookie2=1db157d62a47d1600098fae31701049f; tkSid=1760177613326_556240329_0.0; _samesite_flag_=true; _tb_token_=eaf4a76971871; thw=cn; sdkSilent=1773670236484; havana_sdkSilent=1773670236484; mt=ci=0_0; tkmb=e=okcO90WRuM9wv7nGTZ2_-rX3DTSVrV-TdPW7X9DGPergabDxBccRKcvcnO8iiwnXiiFPB_IQFOVWLajJKI63f9Z9rtbmhnxUteAIBJBhSwPYLuVcajQge_SdNqaVWClOIyCYESgjl5dSVi5Q0yOZYXsbgVor69P7hI1aZ-3ZD_ykjrecxAMSRzo84ttZ8tvjQ9C_kswJLB1Wmw4VzoKVGYqmY1JsNQIob6uvwItC6OcwEmTz_MnbQUcI4VbMweAAe4wMclroAkFVbHpdHLbehIuSgxhL99vOtDuIKwcMY0VKEw17FpewGr-9tCpAtRQVRFcP7nDXPACqoPhMO46pXbiq_Nw6iOfVmQEnT2zBn7-ybhShgiPNswthg69zVkf5DcbShJUCG-rx0nrjo33587ZzzLDMP5b_KdSAomdgdXdgYh-m7jOlAx_awAQ-_g97lGpcJ2moA8OYVhifpb4Tzl3erkpkUHy64oAlkR1vQ3A0pJThyex4uFbZd5s_ymM0NgWDEEGi-3w&iv=0&et=1774000790&tk_cps_param=874030133; XSRF-TOKEN=1b912238-b177-4f14-a0e0-b8abf9aefc02; xlly_s=1; wk_unb=UUphzWModBWn%2FYeb%2BQ%3D%3D; cna=1/pwIGSZkE0CASRwzCE/RY3I; 3PcFlag=1774342155568; sgcookie=E100ycQVgfbIFp6335gwaq%2BO%2FxEdH0XlsCyk5agewE9BE%2BxhJMhzREgh33h86AQGkfMa5hYeBbbLKev5Jubz35B0%2Fvox2mtRIyh23CUCSbgONpw%3D; unb=2207447452491; csg=85e4cf3a; lgc=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; cancelledSubSites=empty; cookie17=UUphzWModBWn%2FYeb%2BQ%3D%3D; dnk=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; skt=3ceb89e630344be3; tracknick=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; _l_g_=Ug%3D%3D; sg=%E6%95%B018; _nk_=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; cookie1=BvXjtrOZ5cHkMZPXVjPH0yv0jKITKPPKC7y20aZJsfA%3D; uc1=cookie16=VFC%2FuZ9az08KUQ56dCrZDlbNdA%3D%3D&cookie21=U%2BGCWk%2F7oPGl&cookie14=UoYZbjZPq%2FA2fA%3D%3D&pas=0&cookie15=WqG3DMC9VAQiUQ%3D%3D&existShop=false; sn=; uc3=vt3=F8dD29X92HDRbMSX6GM%3D&lg2=WqG3DMC9VAQiUQ%3D%3D&nk2=qiAr7FspCes%3D&id2=UUphzWModBWn%2FYeb%2BQ%3D%3D; ultraCookieBase=1k6S5%2BcxkgQpZBn6L%2FWa1H2U5%2F6Ud09f6Bv1MJYGaGXg63vfxNFrby%2FC3EBL8M3Gjn1yNggREmPmr5fV14ekDNOYBJNlpdW%2BCVjKMqM9Y1%2F87KV4F0b0826YhadMR3e8hEqn46%2B0x4d9sFI9wBuM0NltOBBBZfjykwa2OOihHv5tlL5x7EO%2BC3F5Am9Nf8ylezVNVPWHNTO0vNqdvW4GDapMbMGhnoXp%2FfHFoyYc7CZDbOZc%2FXvgpYEFAhF3gnhgmNDE9Gm1QakuKcke%2FTNFBLQDmeBxSMz0RzMQlkhOrESf508jyEM997c0VzfFGD2jE8AadD3t%2BJ0Oa; existShop=MTc3NDM0MjE3Nw%3D%3D; uc4=id4=0%40U2grFnyVW3yldGfZWWJnA8tPzs9LwdHx&nk4=0%40qBqkhc2koMYJ6LPmYMQMMRKOCA%3D%3D; _cc_=UIHiLt3xSw%3D%3D; _hvn_lgc_=0; havana_lgc_exp=1805446177437; havana_lgc2_0=eyJoaWQiOjIyMDc0NDc0NTI0OTEsInNnIjoiMjUzYjY1MTFhMWRmY2IyMGZhMGU0NGYxYmFmNDhkZDYiLCJzaXRlIjowLCJ0b2tlbiI6IjF4d2R4bDI4NHB2V0R4M0hOOGw4SEtRIn0; tfstk=giOq-FxCR1x5p2nH8efZUI0S6jCAs1oIICs1SFYGlijmXCawUeKNcVHAGLSwq3SDismAI5-6-EwsksAN_eKF5EMvf35wRHYb11K1_Nxe2N6fk1gwbhtHbO3AfGjwfFvjPXGBkECOsFoIOXt1M74MQoj0f_VlS17DgQ2ckECOs0a7svGeknNTHmCGs45l7NEGsSjizabdqSbgiNqozNIlsSXGjT2lJwFgI1xMr47OqGfGnhfozNIlj1fis56CQD71nqJAmcmqg1WVxEjzTUdPoSsE9gNiDQb2zMYDB5VMaZW2Z8U7mlfMnEtPdKoUg1LDphQHmomVzI8e_pxoNS5wrpANZQmuPivpzC5XiVGcZI-2iT-g5c_Jdw8NdLggLg9MyQWygvaOyITMM91n9oj6FeRN7CntaH8D0pXHioSzllQoG5O93l2NnaQPA4urSirzkXe1cwyTBtEVzMgDIReOn67PA4yLBRBvlaSIu35..; isg=BPHxr-rLPEMqXZ5-T2ocGw2KAHuL3mVQUJT8oNMGPrjW-hJMGy48IAbcHI6cMv2I'
```
1. 源接口中 cookie 的信息，你需要从 AdToolGyxToken.php 中通过代码获取，在这个表中，union_id = dhh, 接口需要的 _csrf 值需要从 jd_appkey 字段获取，接口需要的 cookie 需要从表中的 cookie 值获取
2. 我项目中此逻辑的接口名为：dhh-forminfo 
- 将获取返回的数据全部返回给前端

# 上传素材逻辑
- 上传素材分为两部分
- 从本地文件夹获取素材，通过大航海接口上传本地素材到远程服务

### 需要一个upload接口，前端会上传文件夹，你需要从文件夹中获取所有素材
- 每一个素材，都需要走一遍以下接口素材逻辑
1. 我给你提供了一个获取校验参数的接口，你需要把这个接口代码化, _csrf 和 cookie还是一样从数据库获取
```
    curl --location 'https://dhh.taobao.com/polystar/api/creative/material/osssign?_csrf=1b912238-b177-4f14-a0e0-b8abf9aefc02' \
--header 'accept: application/json, text/plain, */*' \
--header 'accept-language: zh-CN,zh;q=0.9' \
--header 'bx-v: 2.5.36' \
--header 'priority: u=1, i' \
--header 'referer: https://dhh.taobao.com/' \
--header 'sec-ch-ua: "Not(A:Brand";v="8", "Chromium";v="144", "Google Chrome";v="144"' \
--header 'sec-ch-ua-mobile: ?0' \
--header 'sec-ch-ua-platform: "macOS"' \
--header 'sec-fetch-dest: empty' \
--header 'sec-fetch-mode: cors' \
--header 'sec-fetch-site: same-origin' \
--header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36' \
--header 'x-xsrf-token: 1b912238-b177-4f14-a0e0-b8abf9aefc02' \
--header 'Cookie: t=aad3d47a3693faf2be2983e66f9366c7; wk_cookie2=167f0c4d672ac968e0a93b99b0d0321b; cookie2=1db157d62a47d1600098fae31701049f; tkSid=1760177613326_556240329_0.0; _samesite_flag_=true; _tb_token_=eaf4a76971871; thw=cn; sdkSilent=1773670236484; havana_sdkSilent=1773670236484; mt=ci=0_0; tkmb=e=okcO90WRuM9wv7nGTZ2_-rX3DTSVrV-TdPW7X9DGPergabDxBccRKcvcnO8iiwnXiiFPB_IQFOVWLajJKI63f9Z9rtbmhnxUteAIBJBhSwPYLuVcajQge_SdNqaVWClOIyCYESgjl5dSVi5Q0yOZYXsbgVor69P7hI1aZ-3ZD_ykjrecxAMSRzo84ttZ8tvjQ9C_kswJLB1Wmw4VzoKVGYqmY1JsNQIob6uvwItC6OcwEmTz_MnbQUcI4VbMweAAe4wMclroAkFVbHpdHLbehIuSgxhL99vOtDuIKwcMY0VKEw17FpewGr-9tCpAtRQVRFcP7nDXPACqoPhMO46pXbiq_Nw6iOfVmQEnT2zBn7-ybhShgiPNswthg69zVkf5DcbShJUCG-rx0nrjo33587ZzzLDMP5b_KdSAomdgdXdgYh-m7jOlAx_awAQ-_g97lGpcJ2moA8OYVhifpb4Tzl3erkpkUHy64oAlkR1vQ3A0pJThyex4uFbZd5s_ymM0NgWDEEGi-3w&iv=0&et=1774000790&tk_cps_param=874030133; XSRF-TOKEN=1b912238-b177-4f14-a0e0-b8abf9aefc02; xlly_s=1; wk_unb=UUphzWModBWn%2FYeb%2BQ%3D%3D; cna=1/pwIGSZkE0CASRwzCE/RY3I; sgcookie=E100ycQVgfbIFp6335gwaq%2BO%2FxEdH0XlsCyk5agewE9BE%2BxhJMhzREgh33h86AQGkfMa5hYeBbbLKev5Jubz35B0%2Fvox2mtRIyh23CUCSbgONpw%3D; unb=2207447452491; csg=85e4cf3a; lgc=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; cancelledSubSites=empty; cookie17=UUphzWModBWn%2FYeb%2BQ%3D%3D; dnk=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; skt=3ceb89e630344be3; tracknick=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; _l_g_=Ug%3D%3D; sg=%E6%95%B018; _nk_=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; cookie1=BvXjtrOZ5cHkMZPXVjPH0yv0jKITKPPKC7y20aZJsfA%3D; uc1=cookie16=VFC%2FuZ9az08KUQ56dCrZDlbNdA%3D%3D&cookie21=U%2BGCWk%2F7oPGl&cookie14=UoYZbjZPq%2FA2fA%3D%3D&pas=0&cookie15=WqG3DMC9VAQiUQ%3D%3D&existShop=false; sn=; uc3=vt3=F8dD29X92HDRbMSX6GM%3D&lg2=WqG3DMC9VAQiUQ%3D%3D&nk2=qiAr7FspCes%3D&id2=UUphzWModBWn%2FYeb%2BQ%3D%3D; ultraCookieBase=1k6S5%2BcxkgQpZBn6L%2FWa1H2U5%2F6Ud09f6Bv1MJYGaGXg63vfxNFrby%2FC3EBL8M3Gjn1yNggREmPmr5fV14ekDNOYBJNlpdW%2BCVjKMqM9Y1%2F87KV4F0b0826YhadMR3e8hEqn46%2B0x4d9sFI9wBuM0NltOBBBZfjykwa2OOihHv5tlL5x7EO%2BC3F5Am9Nf8ylezVNVPWHNTO0vNqdvW4GDapMbMGhnoXp%2FfHFoyYc7CZDbOZc%2FXvgpYEFAhF3gnhgmNDE9Gm1QakuKcke%2FTNFBLQDmeBxSMz0RzMQlkhOrESf508jyEM997c0VzfFGD2jE8AadD3t%2BJ0Oa; existShop=MTc3NDM0MjE3Nw%3D%3D; uc4=id4=0%40U2grFnyVW3yldGfZWWJnA8tPzs9LwdHx&nk4=0%40qBqkhc2koMYJ6LPmYMQMMRKOCA%3D%3D; _cc_=UIHiLt3xSw%3D%3D; _hvn_lgc_=0; havana_lgc_exp=1805446177437; havana_lgc2_0=eyJoaWQiOjIyMDc0NDc0NTI0OTEsInNnIjoiMjUzYjY1MTFhMWRmY2IyMGZhMGU0NGYxYmFmNDhkZDYiLCJzaXRlIjowLCJ0b2tlbiI6IjF4d2R4bDI4NHB2V0R4M0hOOGw4SEtRIn0; 3PcFlag=1774422701865; fastSlient=1774422701870; isg=BGdnSuIewr0PMUgYJdgyVUd49p0x7DvOygZKhjnU0fZYKIfqQbw8HqNrSyi2xhNG; tfstk=gYnr7SG_CkcXfuAiQAqeu0VOJRE8vkR6-DNQKvD3F7Vk9D6EuAhUNpL8VqPEijPoquA8-MlIIY_BeumUYAhadYT-OjrECfD5AkhQYJcZMJa7ek9ETXGiTyp8OWVEOv0WhhtseYELxvR6fhNO5MvK4WfQxuE0xCBoYhtseYqLxCO6f4N9JMw3tk4u-SX0MRQuxybk3S2UBwfotkv23RyNZ_j3Enj0BSV3xkchi-VYKk2otkvqnSensVz_8CyQrLuRUUgXt_4YsYVVxGXTUzD6XSSHRmyrFjMlSMjn08zrmZBfZ6qnrYG454RN4kHoWXeiZ_Az30kZYVcDcgrEnVmUmmAch70t3DrSqptum0lrqrlhdBwx5RkU5q9h7SgnHmzq4F6LH0Mn2PZMX_VIlAoULDdp0fkoaV4iq_SPP6eDVMiKz6bUr-e41KJ2K75VehQQNR7dJzCz359o-wQLryy41KJNJwU-C-P6EXf..'
 ```
-  接口返回的数据格式是， 这个返回在之后命名为：osssignResponse
 ```
{"traceId":"215041f717744228327307529eb792","code":0,"data":{"accessid":"LTAIs4hm3acIyRwh","signature":"Wup+8DfWzWROan4AfoWZ6/HEV0s=","expire":"1774423132","host":"https://qihang-material.oss-cn-beijing.aliyuncs.com","callback":null,"dir":"dsp","policy":"eyJleHBpcmF0aW9uIjoiMjAyNi0wMy0yNVQwNzoxODo1Mi43NDRaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0sWyJzdGFydHMtd2l0aCIsIiRrZXkiLCJkc3AiXV19"},"changeFrees":null,"errorCode":null,"changeFree":{"applyOrderURL":""},"message":"","successful":true}
 ```

 2. 获取到 osssignResponse 之后，需要对素材进行接口上传，上传接口如下
 ```
curl --location 'https://qihang-material.oss-cn-beijing.aliyuncs.com/' \
--header 'Accept: */*' \
--header 'Accept-Language: zh-CN,zh;q=0.9' \
--header 'Connection: keep-alive' \
--header 'Origin: https://dhh.taobao.com' \
--header 'Referer: https://dhh.taobao.com/' \
--header 'Sec-Fetch-Dest: empty' \
--header 'Sec-Fetch-Mode: cors' \
--header 'Sec-Fetch-Site: cross-site' \
--header 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36' \
--header 'sec-ch-ua: "Not(A:Brand";v="8", "Chromium";v="144", "Google Chrome";v="144"' \
--header 'sec-ch-ua-mobile: ?0' \
--header 'sec-ch-ua-platform: "macOS"' \
--form 'name="京东-拖把-美数-bjx-0130-22.mp4"' \
--form 'key="dsp/video/WTJYyzwKpZp5pm2fbbmih2i8wR3EijATCMY.mp4"' \
--form 'dir="dsp"' \
--form 'policy="eyJleHBpcmF0aW9uIjoiMjAyNi0wMy0yNVQwNzoxODo1Mi43NDRaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwMF0sWyJzdGFydHMtd2l0aCIsIiRrZXkiLCJkc3AiXV19"' \
--form 'OSSAccessKeyId="LTAIs4hm3acIyRwh"' \
--form 'success_action_status="200"' \
--form 'signature="Wup+8DfWzWROan4AfoWZ6/HEV0s="' \
--form 'file=@"/Users/chemanyu/Desktop/jingcheng/京东-拖把-美数-bjx-0130-22.mp4"'
 ```
- 将这个接口代码化，传递的参数中，
   * name  就是素材名称
   * key = dsp/video/*.mp4，* 是一个 32 位的字符数字随机数，*.mp4 是 素材名，此字段在之后的步骤会用到
   * dir 固定为 dsp
   * policy 是 osssignResponse.policy 的值
   * OSSAccessKeyId 是 osssignResponse.OSSAccessKeyId 的值
   * success_action_status 固定为 200
   * signature 是 osssignResponse.signature 的值
   * file 是保存记录的素材位置
接口返回的只要 httpStatusCode = 200 ,就算上传成功

3. 保存素材数据到待上传列表
在 第二步将素材上传之后，需要通过接口获取素材链接数据，接口如下
```
curl --location 'https://dhh.taobao.com/polystar/api/creative/material/meta/add' \
--header 'accept: */*' \
--header 'accept-language: zh-CN,zh;q=0.9' \
--header 'bx-v: 2.5.36' \
--header 'content-type: application/x-www-form-urlencoded' \
--header 'origin: https://dhh.taobao.com' \
--header 'priority: u=1, i' \
--header 'referer: https://dhh.taobao.com/' \
--header 'sec-ch-ua: "Not(A:Brand";v="8", "Chromium";v="144", "Google Chrome";v="144"' \
--header 'sec-ch-ua-mobile: ?0' \
--header 'sec-ch-ua-platform: "macOS"' \
--header 'sec-fetch-dest: empty' \
--header 'sec-fetch-mode: cors' \
--header 'sec-fetch-site: same-origin' \
--header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36' \
--header 'x-xsrf-token: 1b912238-b177-4f14-a0e0-b8abf9aefc02' \
--header 'Cookie: t=aad3d47a3693faf2be2983e66f9366c7; wk_cookie2=167f0c4d672ac968e0a93b99b0d0321b; cookie2=1db157d62a47d1600098fae31701049f; tkSid=1760177613326_556240329_0.0; _samesite_flag_=true; _tb_token_=eaf4a76971871; thw=cn; sdkSilent=1773670236484; havana_sdkSilent=1773670236484; mt=ci=0_0; tkmb=e=okcO90WRuM9wv7nGTZ2_-rX3DTSVrV-TdPW7X9DGPergabDxBccRKcvcnO8iiwnXiiFPB_IQFOVWLajJKI63f9Z9rtbmhnxUteAIBJBhSwPYLuVcajQge_SdNqaVWClOIyCYESgjl5dSVi5Q0yOZYXsbgVor69P7hI1aZ-3ZD_ykjrecxAMSRzo84ttZ8tvjQ9C_kswJLB1Wmw4VzoKVGYqmY1JsNQIob6uvwItC6OcwEmTz_MnbQUcI4VbMweAAe4wMclroAkFVbHpdHLbehIuSgxhL99vOtDuIKwcMY0VKEw17FpewGr-9tCpAtRQVRFcP7nDXPACqoPhMO46pXbiq_Nw6iOfVmQEnT2zBn7-ybhShgiPNswthg69zVkf5DcbShJUCG-rx0nrjo33587ZzzLDMP5b_KdSAomdgdXdgYh-m7jOlAx_awAQ-_g97lGpcJ2moA8OYVhifpb4Tzl3erkpkUHy64oAlkR1vQ3A0pJThyex4uFbZd5s_ymM0NgWDEEGi-3w&iv=0&et=1774000790&tk_cps_param=874030133; XSRF-TOKEN=1b912238-b177-4f14-a0e0-b8abf9aefc02; xlly_s=1; wk_unb=UUphzWModBWn%2FYeb%2BQ%3D%3D; cna=1/pwIGSZkE0CASRwzCE/RY3I; sgcookie=E100ycQVgfbIFp6335gwaq%2BO%2FxEdH0XlsCyk5agewE9BE%2BxhJMhzREgh33h86AQGkfMa5hYeBbbLKev5Jubz35B0%2Fvox2mtRIyh23CUCSbgONpw%3D; unb=2207447452491; csg=85e4cf3a; lgc=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; cancelledSubSites=empty; cookie17=UUphzWModBWn%2FYeb%2BQ%3D%3D; dnk=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; skt=3ceb89e630344be3; tracknick=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; _l_g_=Ug%3D%3D; sg=%E6%95%B018; _nk_=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; cookie1=BvXjtrOZ5cHkMZPXVjPH0yv0jKITKPPKC7y20aZJsfA%3D; uc1=cookie16=VFC%2FuZ9az08KUQ56dCrZDlbNdA%3D%3D&cookie21=U%2BGCWk%2F7oPGl&cookie14=UoYZbjZPq%2FA2fA%3D%3D&pas=0&cookie15=WqG3DMC9VAQiUQ%3D%3D&existShop=false; sn=; uc3=vt3=F8dD29X92HDRbMSX6GM%3D&lg2=WqG3DMC9VAQiUQ%3D%3D&nk2=qiAr7FspCes%3D&id2=UUphzWModBWn%2FYeb%2BQ%3D%3D; ultraCookieBase=1k6S5%2BcxkgQpZBn6L%2FWa1H2U5%2F6Ud09f6Bv1MJYGaGXg63vfxNFrby%2FC3EBL8M3Gjn1yNggREmPmr5fV14ekDNOYBJNlpdW%2BCVjKMqM9Y1%2F87KV4F0b0826YhadMR3e8hEqn46%2B0x4d9sFI9wBuM0NltOBBBZfjykwa2OOihHv5tlL5x7EO%2BC3F5Am9Nf8ylezVNVPWHNTO0vNqdvW4GDapMbMGhnoXp%2FfHFoyYc7CZDbOZc%2FXvgpYEFAhF3gnhgmNDE9Gm1QakuKcke%2FTNFBLQDmeBxSMz0RzMQlkhOrESf508jyEM997c0VzfFGD2jE8AadD3t%2BJ0Oa; existShop=MTc3NDM0MjE3Nw%3D%3D; uc4=id4=0%40U2grFnyVW3yldGfZWWJnA8tPzs9LwdHx&nk4=0%40qBqkhc2koMYJ6LPmYMQMMRKOCA%3D%3D; _cc_=UIHiLt3xSw%3D%3D; _hvn_lgc_=0; havana_lgc_exp=1805446177437; havana_lgc2_0=eyJoaWQiOjIyMDc0NDc0NTI0OTEsInNnIjoiMjUzYjY1MTFhMWRmY2IyMGZhMGU0NGYxYmFmNDhkZDYiLCJzaXRlIjowLCJ0b2tlbiI6IjF4d2R4bDI4NHB2V0R4M0hOOGw4SEtRIn0; 3PcFlag=1774422701865; fastSlient=1774422701870; isg=BGdnSuIewr0PMUgYJdgyVUd49p0x7DvOygZKhjnU0fZYKIfqQbw8HqNrSyi2xhNG; tfstk=gYnr7SG_CkcXfuAiQAqeu0VOJRE8vkR6-DNQKvD3F7Vk9D6EuAhUNpL8VqPEijPoquA8-MlIIY_BeumUYAhadYT-OjrECfD5AkhQYJcZMJa7ek9ETXGiTyp8OWVEOv0WhhtseYELxvR6fhNO5MvK4WfQxuE0xCBoYhtseYqLxCO6f4N9JMw3tk4u-SX0MRQuxybk3S2UBwfotkv23RyNZ_j3Enj0BSV3xkchi-VYKk2otkvqnSensVz_8CyQrLuRUUgXt_4YsYVVxGXTUzD6XSSHRmyrFjMlSMjn08zrmZBfZ6qnrYG454RN4kHoWXeiZ_Az30kZYVcDcgrEnVmUmmAch70t3DrSqptum0lrqrlhdBwx5RkU5q9h7SgnHmzq4F6LH0Mn2PZMX_VIlAoULDdp0fkoaV4iq_SPP6eDVMiKz6bUr-e41KJ2K75VehQQNR7dJzCz359o-wQLryy41KJNJwU-C-P6EXf..' \
--data-urlencode 'ossUrl=https://qihang-material.oss-cn-beijing.aliyuncs.com/dsp/video/WTJYyzwKpZp5pm2fbbmih2i8wR3EijATCMY.mp4' \
--data-urlencode 'fileName=京东-拖把-美数-bjx-0130-22.mp4' \
--data-urlencode '_csrf=1b912238-b177-4f14-a0e0-b8abf9aefc02'
```
- 返回的数据结果为, 此结果保存为 metaResponse
```
{"traceId":"2150424317744229700504558eba7d","code":0,"data":{"materialUrl":"https://qh-material.taobao.com/dsp/video/WTJYyzwKpZp5pm2fbbmih2i8wR3EijATCMY.mp4","gmtModified":null,"fileName":"京东-拖把-美数-bjx-0130-22.mp4","materialType":2,"accountType":2,"creatorId":2207447452491,"materialCode":"12ef012a239544b181842eda4e78c054","duplicate":false,"gmtCreate":null,"posterUrl":"https://qh-material.taobao.com/video/image/12ef012a239544b181842eda4e78c054.jpg","name":null,"id":8591660,"materialSource":2,"status":1},"changeFrees":null,"errorCode":null,"changeFree":{"applyOrderURL":""},"message":"","successful":true}
```

- 将这个接口逻辑代码化
 * _csrf， 和 cookie 还是一样从数据库获取的
 * 返回的数据是 metaResponse
 * 调用接口的时候 参数 fileName 就是素材名称
 * 参数 ossUrl 是 osssignResponse.host/步骤二中的 key 值

4. 以上三步骤，是上传素材的逻辑，每一个素材，都会得到一个 metaResponse，最后我希望这个上传素材的接口能够返回一个 metaList, 这个list里是所有素材的 metaResponse.fileName, metaResponse.materialUrl, metaResponse.materialCode. 用于前端展示

### 提交审核接口
- 需要一个 materialCreate 接口
- 上一步骤，返回了 metaList 用于前端展示，然后前端会把这个数据给到此接口，你来定义接收格式，之后同步前端
- 以下是这个提交素材的源接口逻辑，帮我代码化
```
curl --location 'https://dhh.taobao.com/polystar/api/creative/material/create' \
--header 'accept: */*' \
--header 'accept-language: zh-CN,zh;q=0.9' \
--header 'bx-v: 2.5.36' \
--header 'content-type: application/x-www-form-urlencoded' \
--header 'origin: https://dhh.taobao.com' \
--header 'priority: u=1, i' \
--header 'referer: https://dhh.taobao.com/' \
--header 'sec-ch-ua: "Not(A:Brand";v="8", "Chromium";v="144", "Google Chrome";v="144"' \
--header 'sec-ch-ua-mobile: ?0' \
--header 'sec-ch-ua-platform: "macOS"' \
--header 'sec-fetch-dest: empty' \
--header 'sec-fetch-mode: cors' \
--header 'sec-fetch-site: same-origin' \
--header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36' \
--header 'x-xsrf-token: 1b912238-b177-4f14-a0e0-b8abf9aefc02' \
--header 'Cookie: t=aad3d47a3693faf2be2983e66f9366c7; wk_cookie2=167f0c4d672ac968e0a93b99b0d0321b; cookie2=1db157d62a47d1600098fae31701049f; tkSid=1760177613326_556240329_0.0; _samesite_flag_=true; _tb_token_=eaf4a76971871; thw=cn; sdkSilent=1773670236484; havana_sdkSilent=1773670236484; mt=ci=0_0; tkmb=e=okcO90WRuM9wv7nGTZ2_-rX3DTSVrV-TdPW7X9DGPergabDxBccRKcvcnO8iiwnXiiFPB_IQFOVWLajJKI63f9Z9rtbmhnxUteAIBJBhSwPYLuVcajQge_SdNqaVWClOIyCYESgjl5dSVi5Q0yOZYXsbgVor69P7hI1aZ-3ZD_ykjrecxAMSRzo84ttZ8tvjQ9C_kswJLB1Wmw4VzoKVGYqmY1JsNQIob6uvwItC6OcwEmTz_MnbQUcI4VbMweAAe4wMclroAkFVbHpdHLbehIuSgxhL99vOtDuIKwcMY0VKEw17FpewGr-9tCpAtRQVRFcP7nDXPACqoPhMO46pXbiq_Nw6iOfVmQEnT2zBn7-ybhShgiPNswthg69zVkf5DcbShJUCG-rx0nrjo33587ZzzLDMP5b_KdSAomdgdXdgYh-m7jOlAx_awAQ-_g97lGpcJ2moA8OYVhifpb4Tzl3erkpkUHy64oAlkR1vQ3A0pJThyex4uFbZd5s_ymM0NgWDEEGi-3w&iv=0&et=1774000790&tk_cps_param=874030133; XSRF-TOKEN=1b912238-b177-4f14-a0e0-b8abf9aefc02; xlly_s=1; wk_unb=UUphzWModBWn%2FYeb%2BQ%3D%3D; cna=1/pwIGSZkE0CASRwzCE/RY3I; sgcookie=E100ycQVgfbIFp6335gwaq%2BO%2FxEdH0XlsCyk5agewE9BE%2BxhJMhzREgh33h86AQGkfMa5hYeBbbLKev5Jubz35B0%2Fvox2mtRIyh23CUCSbgONpw%3D; unb=2207447452491; csg=85e4cf3a; lgc=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; cancelledSubSites=empty; cookie17=UUphzWModBWn%2FYeb%2BQ%3D%3D; dnk=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; skt=3ceb89e630344be3; tracknick=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; _l_g_=Ug%3D%3D; sg=%E6%95%B018; _nk_=%5Cu4E0A%5Cu6D77%5Cu7F8E%5Cu6570; cookie1=BvXjtrOZ5cHkMZPXVjPH0yv0jKITKPPKC7y20aZJsfA%3D; uc1=cookie16=VFC%2FuZ9az08KUQ56dCrZDlbNdA%3D%3D&cookie21=U%2BGCWk%2F7oPGl&cookie14=UoYZbjZPq%2FA2fA%3D%3D&pas=0&cookie15=WqG3DMC9VAQiUQ%3D%3D&existShop=false; sn=; uc3=vt3=F8dD29X92HDRbMSX6GM%3D&lg2=WqG3DMC9VAQiUQ%3D%3D&nk2=qiAr7FspCes%3D&id2=UUphzWModBWn%2FYeb%2BQ%3D%3D; ultraCookieBase=1k6S5%2BcxkgQpZBn6L%2FWa1H2U5%2F6Ud09f6Bv1MJYGaGXg63vfxNFrby%2FC3EBL8M3Gjn1yNggREmPmr5fV14ekDNOYBJNlpdW%2BCVjKMqM9Y1%2F87KV4F0b0826YhadMR3e8hEqn46%2B0x4d9sFI9wBuM0NltOBBBZfjykwa2OOihHv5tlL5x7EO%2BC3F5Am9Nf8ylezVNVPWHNTO0vNqdvW4GDapMbMGhnoXp%2FfHFoyYc7CZDbOZc%2FXvgpYEFAhF3gnhgmNDE9Gm1QakuKcke%2FTNFBLQDmeBxSMz0RzMQlkhOrESf508jyEM997c0VzfFGD2jE8AadD3t%2BJ0Oa; existShop=MTc3NDM0MjE3Nw%3D%3D; uc4=id4=0%40U2grFnyVW3yldGfZWWJnA8tPzs9LwdHx&nk4=0%40qBqkhc2koMYJ6LPmYMQMMRKOCA%3D%3D; _cc_=UIHiLt3xSw%3D%3D; _hvn_lgc_=0; havana_lgc_exp=1805446177437; havana_lgc2_0=eyJoaWQiOjIyMDc0NDc0NTI0OTEsInNnIjoiMjUzYjY1MTFhMWRmY2IyMGZhMGU0NGYxYmFmNDhkZDYiLCJzaXRlIjowLCJ0b2tlbiI6IjF4d2R4bDI4NHB2V0R4M0hOOGw4SEtRIn0; 3PcFlag=1774422701865; fastSlient=1774422701870; isg=BGdnSuIewr0PMUgYJdgyVUd49p0x7DvOygZKhjnU0fZYKIfqQbw8HqNrSyi2xhNG; tfstk=gYnr7SG_CkcXfuAiQAqeu0VOJRE8vkR6-DNQKvD3F7Vk9D6EuAhUNpL8VqPEijPoquA8-MlIIY_BeumUYAhadYT-OjrECfD5AkhQYJcZMJa7ek9ETXGiTyp8OWVEOv0WhhtseYELxvR6fhNO5MvK4WfQxuE0xCBoYhtseYqLxCO6f4N9JMw3tk4u-SX0MRQuxybk3S2UBwfotkv23RyNZ_j3Enj0BSV3xkchi-VYKk2otkvqnSensVz_8CyQrLuRUUgXt_4YsYVVxGXTUzD6XSSHRmyrFjMlSMjn08zrmZBfZ6qnrYG454RN4kHoWXeiZ_Az30kZYVcDcgrEnVmUmmAch70t3DrSqptum0lrqrlhdBwx5RkU5q9h7SgnHmzq4F6LH0Mn2PZMX_VIlAoULDdp0fkoaV4iq_SPP6eDVMiKz6bUr-e41KJ2K75VehQQNR7dJzCz359o-wQLryy41KJNJwU-C-P6EXf..' \
--data-urlencode 'admissionType=material' \
--data-urlencode 'appId=1' \
--data-urlencode 'taskType=400' \
--data-urlencode 'adType=1' \
--data-urlencode 'scenarioType=2' \
--data-urlencode 'hotEvent=0' \
--data-urlencode 'baseImageType=0' \
--data-urlencode 'customTitle=测试' \
--data-urlencode 'customCopy=测试' \
--data-urlencode 'bizType=pool_13129' \
--data-urlencode 'scenarioTypeDesc=' \
--data-urlencode 'bizTypeDesc=' \
--data-urlencode 'materialList=[{"materialCode":"12ef012a239544b181842eda4e78c054","fileId":"o_1jkhor4jhmv4buh1uu1ir01q19cmy","materialUrl":"https://qh-material.taobao.com/dsp/video/WTJYyzwKpZp5pm2fbbmih2i8wR3EijATCMY.mp4"}]' \
--data-urlencode '_csrf=1b912238-b177-4f14-a0e0-b8abf9aefc02'

```
- 返回的数据格式
```
{"traceId":"2150421217744230670455886e400f","code":0,"data":null,"changeFrees":null,"errorCode":null,"changeFree":{"applyOrderURL":""},"message":"","successful":true}
```
 * 前端传递进来的metaList数据，都要用这个逻辑上传过去
 * 参数中 materialList 是一个传递列表，最多只能有十个素材
 * materialCode 就是 metaResponse.materialCode
 * materialUrl 就是 metaResponse.materialUrl
 * fileId 应该是个随机数，你按照这个格式随机生成

- 循环将所有素材传递，你需要记录成功数据和失败数据，将此数据返回给前端


## 结尾
- 现阶段，代码逻辑最好加上日志，用于之后联调
- 写完接口之后，将这些接口写一个 md 文档，方便我给到前端项目联调。文档中需要有 curl 的测试用例