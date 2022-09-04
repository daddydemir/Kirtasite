# Kırtasite
Bu proje sayesinde kırtasiyelerde sıra beklemenize gerek kalmaz, siparişinizi verirsiniz ve siparişiniz hazırlandığında size haber veririz. 
Bu repo **Kırtasite** uygulamasının backend kodlarını içermektedir.

---
Base Url
```
api/v1/
```
<details>
	<summary><h3> USER </h3></summary>


```json
{
	"Id":0,
	"RoleId":0,
	"Password": "",
	"ImagePath":"",
	"Phone":"",
	"Mail":"",
	"Role": {}
}
```
Image Upload & Update
```
user/image/0
```
>  İstek **multipart/form-data** headerı ile yapılmalıdır. key => "file".

[ÖRNEK KULLANIM](https://github.com/daddydemir/Kirtasite/blob/main/docs/ImageUpload.md)


Password Update
```
user/password/0
```

</details>


<details>
	<summary><h3> CUSTOMER</h3></summary>
	
```json
{
	"UserId":0,
	"Username":"",
	"User":{}
}
```
Get
```
customers
```
Get
```
customer/0
```
Add 
```
customer
```
Put
```
customer/0
```
</details>


<details>
	<summary><h3> STATIONERY</h3></summary>
	
```json
{
	"UserId":0,
	"AddressId":0,
	"CompanyName":"",
	"Score":0.0,
	"User":{},
	"Address":{}
}
```
Get
```
stationeries
```
Get 
```
stationery/0
```
Post
```
stationery
```
Put
```
stationery/0
```
</details>


<details>
	<summary><h3> ADDRESS</h3></summary>
	
```json
{
	"Id":0,
	"CityId":0,
	"DistrictId":0,
	"Header":"",
	"X":"",
	"Y":"",
	"City":{},
	"District":{}
}
```
Get 
```
address/0
```
Get By CityId
```
addresses/city/0
```
Post
```
address
```
Put
```
address/0
```
</details>


<details>
	<summary><h3> FILE</h3></summary>
	
> Dosya eklemek için **multipart/form-data** headeri kullanılmalıdır. key => "file"
	
[ÖRNEK KULLANIM](https://github.com/daddydemir/Kirtasite/blob/main/docs/FileUpload.md)
	
```json
{
	"Id":0,
	"CustomerId":0,
	"Private":true,
	"FilePath":"",
	"FolderId":"",
	"CreatedDate":'00-00-0000',
	"Customer":{}
}
```
Get
```
file/0
```
Get By CustomerId
```
files/customer/0
```
Post
```
file
```
Put
```
file/0
```
</details>


<details>
	<summary><h3> PRICE</h3></summary>
	
```json
{
	"Id":0,
	"StationeryId":0,
	"Info":"",
	"Price":0.0
}
```
Get
```
prices
```
Get
```
price/0
```
Get By StationeryId
```
prices/stationery/0
```
Post
```
price
```
Put
```
price/0
```
</details>


<details>
	<summary><h3> ORDER </h3></summary>

```json
{
	"Id":0,
	"FileId":0,
	"CustomerId":0,
	"StationeryId":0,
	"StatusId":0,
	"TotalPrice":0.0,
	"CreatedDate":'00-00-0000',
	"DeliveryDate":'00-00-0000',
	"File":{},
	"Customer":{},
	"Stationery":{},
	"Status":{}
}
```
Get
```
order/0
```
Get By CustomerId
```
orders/customer/0
```
Get By StationeryId
```
orders/stationery/0
```
Post
```
order
```
Put Cancel
```
order/0/cancel
```
Put Confirm
```
order/0/confirm
```
Put Ready
```
order/0/ready
```
Put Complete
```
order/0/complete
```

</details>


<details>
	<summary><h3> STATUS </h3></summary>

```json
{
	"Id":0,
	"Content":""
}
```
Get
```
statuses
```
Get
```
status/0
```

</details>


<details>
	<summary><h3> COMMENT </h3></summary>

```json
{
	"Id":0,
	"CustomerId":0,
	"StationeryId":0,
	"Content":"",
	"CreatedDate":'00-00-0000',
	"Score":0.0,
	"Customer":{},
	"Stationery":{}
}
```
Get 
```
comment/0
```
Get By StationeryId
```
comments/stationery/0
```
Get By CustomerId
```
comments/customer/0
```
Post
```
comment
```
Put
```
comment/0
```

</details>
