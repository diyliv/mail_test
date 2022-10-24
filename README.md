# Тестовое задание 

## Сборка
``` make build ```

## Роут для отправки рассылок 

| Property      | Value         |
| ------------- | ------------- |
| API           | /send_email      |
| Method        | 'POST'        |
| Description   | Данный эндпоинт отвечает за отправку рассылок |

Body:
```json
{
    "sender_email":"from@gmail.com", // кто будет отправлять
    "sender_msg":"hello", // отправляемое письмо
    "recipients_email":["first_recipient@gmail.com", "second_recipient@gmail.com"] // получатели
}
```

## Роут для отправки отложенных рассылок

| Property      | Value         |
| ------------- | ------------- |
| API           | /delay_email      |
| Method        | 'POST'        |
| Description   | Данный эндпоинт отвечает за отправку отложенных рассылок |

Body:
```json
{
    "sender_email":"from@gmail.com",
    "sender_msg":"hello",
    "recipients_email":["first@gmail.com", "second@gmail.com"]
}
```

## База 
Вот что лежит в базе по итогу совершенных выше операций
```
 sender_email: from@gmail.com 
 sender_msg:  hello       
 sender_password: 1234            
 sender_hashed_pass: $2a$10$izVFkmJJGbLD56oBoRxGh.51xkmDsFU0F8L7kC5ve2rM/GiNUDQK6 
 recipients_email: {first@gmail.com,second@gmail.com} 
 email_status: sent          
 sent_at: 2022-10-24 22:15:25.774174

```

## Улучшения
Создал бы воркеров, чтобы бегали и отправляли сообщения(worker pool)
Сделал бы Dockerfile поменьше 