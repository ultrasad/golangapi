package models

//Customer return customer detail
type Customer struct {
	Active        string `json:"active"`
	Age           string `json:"age"`
	BilltoAddress struct {
		Address          string `json:"address"`
		AddressID        string `json:"address_id"`
		AddressType      string `json:"address_type"`
		Amphur           string `json:"amphur"`
		Block            string `json:"block"`
		CustomerEmail    string `json:"customer_email"`
		CustomerFullname string `json:"customer_fullname"`
		CustomerSms      string `json:"customer_sms"`
		CustomerTaxID    string `json:"customer_tax_id"`
		DefaultAddress   string `json:"default_address"`
		District         string `json:"district"`
		EasyTel          string `json:"easy_tel"`
		Province         string `json:"province"`
		Street           string `json:"street"`
		Zipcode          string `json:"zipcode"`
	} `json:"billto_address"`
	Birthdate           string `json:"birthdate"`
	ContactName         string `json:"contact_name"`
	ContactNameEn       string `json:"contact_name_en"`
	ContactTaxid        string `json:"contact_taxid"`
	CreatedBy           string `json:"created_by"`
	CustomerEmail       string `json:"customer_email"`
	CustomerGroup       string `json:"customer_group"`
	CustomerMasterGroup string `json:"customer_master_group"`
	CustomerMobile      string `json:"customer_mobile"`
	CustomerName        string `json:"customer_name"`
	CustomerNameEn      string `json:"customer_name_en"`
	CustomerRefno       string `json:"customer_refno"`
	CustomerResend      string `json:"customer_resend"`
	CustomerTaxid       string `json:"customer_taxid"`
	CustomerType        string `json:"customer_type"`
	DataResend          string `json:"data_resend"`
	EasyTel             string `json:"easy_tel"`
	EsUpdate            string `json:"es_update"`
	Fax                 string `json:"fax"`
	Firstname           string `json:"firstname"`
	FirstnameEn         string `json:"firstname_en"`
	FullBilltoAddress   string `json:"full_billto_address"`
	FullShiptoAddress   string `json:"full_shipto_address"`
	InsertDatetime      string `json:"insert_datetime"`
	Lastname            string `json:"lastname"`
	LastnameEn          string `json:"lastname_en"`
	MarriageStatus      string `json:"marriage_status"`
	MemberChannel       string `json:"member_channel"`
	MemberStatus        string `json:"member_status"`
	MobileOperator      string `json:"mobile_operator"`
	ModifiedBy          string `json:"modified_by"`
	Nickname            string `json:"nickname"`
	OfflineCustomer     string `json:"offline_customer"`
	Remark1             string `json:"remark1"`
	Remark10            string `json:"remark10"`
	Remark11            string `json:"remark11"`
	Remark12            string `json:"remark12"`
	Remark13            string `json:"remark13"`
	Remark14            string `json:"remark14"`
	Remark15            string `json:"remark15"`
	Remark2             string `json:"remark2"`
	Remark3             string `json:"remark3"`
	Remark4             string `json:"remark4"`
	Remark5             string `json:"remark5"`
	Remark6             string `json:"remark6"`
	Remark7             string `json:"remark7"`
	Remark8             string `json:"remark8"`
	Remark9             string `json:"remark9"`
	Sex                 string `json:"sex"`
	ShiptoAddress       struct {
		Address          string `json:"address"`
		AddressID        string `json:"address_id"`
		AddressType      string `json:"address_type"`
		Amphur           string `json:"amphur"`
		Block            string `json:"block"`
		CustomerEmail    string `json:"customer_email"`
		CustomerFullname string `json:"customer_fullname"`
		CustomerSms      string `json:"customer_sms"`
		CustomerTaxID    string `json:"customer_tax_id"`
		DefaultAddress   string `json:"default_address"`
		District         string `json:"district"`
		EasyTel          string `json:"easy_tel"`
		Province         string `json:"province"`
		Street           string `json:"street"`
		Zipcode          string `json:"zipcode"`
	} `json:"shipto_address"`
	SourceSystem   string `json:"source_system"`
	Telephone      string `json:"telephone"`
	UpdateDatetime string `json:"update_datetime"`
}

//Customers is Customer
type Customers struct {
	Customers []Customer
}
