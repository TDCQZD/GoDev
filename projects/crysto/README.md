# CRYSTO

## 社区机器人
功能：爬取数据，机器人管理

## Telegram 机器人
### 创建机器人
1. https://telegram.me/botfather,打开botfarther的聊天对话框
2. 创建一个bot
    - /newbot
    - 输入你要创建的bot名字
    - 创建的bot名字(TetrisBot or tetris_bot)
    > 对话框输入内容回车发送
3. 获取api token
### 机器人添加到group
- 搜索bot name，添加

    点击 SEND MESSAGE    
- https://api.telegram.org/botxxx:xxx/getUpdates 
    
    获取json 中id
### 机器人发送信息
```
curl -X POST "https://api.telegram.org/botXXX:YYYY/sendMessage" -d "chat_id=-zzzzzzzzzz&text=my sample text"

https://api.telegram.org/botXXX:YYYY/sendMessage?chat_id=-zzzzzzzzzz&text=my sample text
```
- XXX:YYYY : bot+token，一定要加上bot前缀
- chat_id : group ID
- text : 发送内容
### API
* https://api.telegram.org/bot815628366:AAFH7P-dZernTDrDMYl0J7aT8wNa7900-5M/getMe
* https://api.telegram.org/bot891757599:AAG-mnzXoMz-jaGiUW7t9dYhUJa0_0lpRiE/getUpdates
* https://api.telegram.org/bot891757599:AAG-mnzXoMz-jaGiUW7t9dYhUJa0_0lpRiE/sendMessage?chat_id=-282646467&text="Hello Hot" 
### 参考文档
* https://core.telegram.org/bots/api
* https://fullmeter.com/blog/?p=14
* http://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id-ruby-gem-telegram-bot
* https://core.telegram.org/bots
* https://core.telegram.org/bots/api