# format_definition
format_definition はデータ連携基盤におけるオーダーAPIの仕様を書き記したものです。

## APIサービス名:【Invoice Document > DPFM_API_INVOICE_DOCUMENT_SRV】
### API名:【Invoice Document Header > A_Header】
### Accepter名:【Header】

### <項目>
以下の表には、data-platform-api-invoice-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                   | Description                                                                                                                                                           | EC  | 
| -------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| InvoiceDocument            | 請求伝票。自動登録される。変更不可。                                                                                                                                  |     | 
| CreationDate               | 作成日付。自動生成される。変更不可。                                                                                                                                  |     | 
| LastChangeDate             | 最終更新日付。自動生成される。変更不可。                                                                                                                              |     | 
| BillToParty                | 請求先。オーダーの取引先機能から参照される。変更不可。                                                                                                                |     | 
| BillFromParty              | 請求元。オーダーの取引先機能から参照される。変更不可。                                                                                                                |     | 
| InvoiceDocumentDate        | 請求伝票日付。請求日を入力する。                                                                                                                                      |     | 
| InvoiceDocumentTime        | 請求伝票時刻。請求時刻を入力する。                                                                                                                                    |     | 
| InvoicePeriodStartDate     | 請求期間開始日付。                                                                                                                                                    |     | 
| InvoicePeriodEndDate       | 請求期間終了日付。                                                                                                                                                    |     | 
| AccountingPostingDate      | 会計転記日付。                                                                                                                                                        |     | 
| InvoiceDocumentIsCancelled | 請求伝票キャンセル済。請求伝票がキャンセルされた場合、trueが設定される。                                                                                              |     | 
| CancelledInvoiceDocument   | キャンセルされた請求伝票。                                                                                                                                            |     | 
| IsExportDelivery           | 輸出入フラグ。請求対象が輸出入取引である場合、trueが設定される。                                                                                                      |     | 
| TotalNetAmount             | 合計正味金額。消費税を除いた合計正味金額が自動計算される。必要に応じて変更する。明細の正味金額の合計との整合性がチェックされる。                                      |     | 
| TotalTaxAmount             | 合計消費税額。合計消費税額が自動計算される。必要に応じて変更する。合計総額と合計正味金額との差から計算される消費税総額の理論値に2円以上の差がある場合、エラーとなる。 |     | 
| TotalGrossAmount           | 合計総額。消費税を含んだ合計総額が自動計算される。必要に応じて変更する。明細の総額の合計との整合性がチェックされる。                                                  |     | 
| TransactionCurrency        | 取引通貨。オーダーの取引通貨が設定される。変更不可。                                                                                                                  |     | 
| Incoterms                  | インコタームズ分類。入出荷伝票またはオーダーのインコタームズが設定される。変更不可。                                                                                  |     | 
| PaymentTerms               | 支払条件。入出荷伝票またはオーダーの支払条件が設定される。必要に応じて変更する。                                                                                      | ✔  | 
| DueCalculationBaseDate     | 期日計算基準日。請求書の締め日のこと。支払条件に基づいて自動設定される。必要に応じて変更する。                                                                        |     | 
| NetPaymentDays             | 正味支払日数。締め日から支払日までの日数のこと。必要に応じて変更する。                                                                                                |     | 
| PaymentMethod              | 支払方法。入出荷伝票またはオーダーの支払方法が設定される。必要に応じて変更する。                                                                                      | ✔  | 
| PaymentBlockingReason      | 支払ブロック理由。支払請求に対して支払をブロックする場合、trueを入力する。                                                                                            |     | 
| ExternalReferenceDocument  | 外部参照伝票。周辺業務システムの参照伝票番号を入力する。                                                                                                              |     | 
| DocumentHeaderText         | 請求書ヘッダテキスト。                                                                                                                                                |     | 

### API名:【Invoice Document Header Partner > A_HeaderPartner】
### Accepter名:【HeaderPartner】

### <項目>
以下の表には、data-platform-api-invoice-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                              | Description                                                                                                                | EC  | 
| ------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | --- | 
| PartnerFunction                       | 取引先機能。入出荷伝票またはオーダーから参照される。変更不可。<br> 参照される取引先機能: <br> BUYER:買い手<br> SELLER:売り手 <br> CUSTOMER:受注先 <br> SUPPLIER:仕入先 <br> MANUFACTURER:製造者 <br> DELIVERFROM:入出荷元 <br> DELIVERTO:入出荷先 <br> LOGI:物流業者 <br> BILLTO:請求先 <br> BILLFROM:請求元 <br> PAYEE:支払人 <br> RECEIVER:受取人 <br> PSPROVIDER:支払決済サービスプロバイダ |     |                                                                                                                       | 
| BusinessPartner                       | ビジネスパートナコード。取引先機能に対応するビジネスパートナコードが設定される。変更不可。                                 |     | 
| BusinessPartnerFullName               | ビジネスパートナフルネーム。オーダーまたは入出荷伝票から設定される。変更不可。                                             |     | 
| BusinessPartnerName                   | ビジネスパートナ名称。オーダーまたは入出荷伝票から設定される。変更不可。                                                   |     | 
| OrderType                             | オーダータイプ。オーダーまたは入出荷伝票から設定される。変更不可。                                                         |     | 
| Organization                          | 販売組織または購買組織。オーダーまたは入出荷伝票から設定される。変更不可。                                                 |     | 
| Language                              | 言語コード。オーダーまたは入出庫伝票から設定される。変更不可。                                                             |     | 
| Currency                              | 通貨コード。オーダーまたは入出庫伝票から設定される。変更不可。                                                             |     | 
| ExternalDocumentID                    | 外部文書ID。                                                                                                               |     | 
| AddressID                             | 住所ID。オーダーまたは入出荷伝票の住所IDが設定される。マニュアルで住所を更新する場合、住所IDが新たに設定される。変更不可。 |     | 

### API名:【Invoice Document Header Partner Contact > A_HeaderPartnerContact】
### Accepter名:【HeaderPartnerContact】

### <項目>
以下の表には、data-platform-api-invoice-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property          | Description                                                                                                                                                                | EC  | 
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --- | 
| PartnerFunction   | 取引先機能。ヘッダパートナから設定される。変更不可。                                                                                                                       |     | 
| ContactID         | コンタクトID。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。       |     | 
| BusinessPartner   | ビジネスパートナコード。ヘッダパートナから設定される。変更不可。                                                                                                           |     | 
| ContactPersonName | コンタクト担当者名。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。 |     | 
| EmailAddress      | Eメールアドレス。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 
| PhoneNumber       | 電話番号。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。           |     | 
| MobilePhoneNumber | モバイル電話番号。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。   |     | 
| FaxNumber         | ファクス番号。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。       |     | 
| ContactTag1       | コンタクトタグ1。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 
| ContactTag2       | コンタクトタグ2。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 
| ContactTag3       | コンタクトタグ3。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 
| ContactTag4       | コンタクトタグ4。ビジネスパートナ得意先取引先機能コンタクトデータまたは仕入先取引先機能コンタクトデータから設定される。必要に応じて変更する。参照の場合は変更できない。    |     | 


### API名:【Invoice Document Header PDF > A_HeaderPDF】
### Accepter名:【HeaderPDF】

### <項目>
以下の表には、data-platform-api-invoice-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。


| Property                 | Description                                                                    | EC  | 
| ------------------------ | ------------------------------------------------------------------------------ | --- | 
| DocType                  | 文書タイプ。請求伝票の場合、次から選択する。 <br> INVOICE:請求書           |     |                                                                           | 
| DocVersionID             | 文書のバージョンID。                                                           |     | 
| DocID                    | 文書のID。ハッシュ値で登録される。変更不可。                                   |     | 
| DocIssuerBusinessPartner | 文書発行者ビジネスパートナコード。ビジネスパートナマスタから選択して入力する。 |     | 
| FileName                 | ファイル名称。                                                                 |     | 


### API名:【Invoice Document Item> A_Item】
### Accepter名:【Item】

### <項目>
以下の表には、data-platform-api-invoice-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                       | Description                                                                                                                                                                          | EC  | 
| ------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | --- | 
| InvoiceDocumentItem            | 請求伝票明細。入伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                         |     | 
| DocumentItemCategory           | 伝票明細カテゴリ。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                 |     | 
| CreationDate                   | 作成日付。自動生成される。変更不可。                                                                                                                                                 |     | 
| CreationTime                   | 作成時刻。自動生成される。変更不可。                                                                                                                                                 |     | 
| BusinessArea                   | 事業領域。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                         |     | 
| DeliverToParty                 | 出荷先または入荷先。入出荷伝票から参照される。変更不可。                                                                                                                             |     | 
| DeliverFromParty               | 出荷元または入荷元。入出荷伝票から参照される。変更不可。                                                                                                                             |     | 
| ProductStandardID              | 国際商品識別コード。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                               |     | 
| Batch                          | ロット番号。入出荷伝票明細から参照して登録される。変更不可。                                                                                                                         |     | 
| Product                        | 品目コード。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                       |     | 
| ProductGroup                   | 品目グループ。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                     |     | 
| ProductionPlantPartnerFunction | 製造プラント取引先機能。入出荷伝票から参照して登録される。変更不可。                                                                                                                 |     | 
| ProductionPlantBusinessPartner | 製造プラントビジネスパートナ。入出荷伝票から参照して登録される。変更不可。                                                                                                           |     | 
| ProductionPlant                | 製造プラント。入出荷伝票から参照して登録される。変更不可。                                                                                                                           |     | 
| ProductionPlantStorageLocation | 製造プラントの保管場所。入出荷伝票から参照して登録される。変更不可。                                                                                                                 |     | 
| IssuingPlantPartnerFunction    | 出荷プラント取引先機能。入出荷伝票から参照して登録される。変更不可。                                                                                                                 |     | 
| IssuingPlantBusinessPartner    | 出荷プラントビジネスパートナ。入出荷伝票から参照して登録される。変更不可。                                                                                                           |     | 
| IssuingPlant                   | 出荷プラント。入出荷伝票から参照して登録される。変更不可。                                                                                                                           |     | 
| IssuingPlantStorageLocation    | 出荷プラントの保管場所。入出荷伝票から参照して登録される。変更不可。                                                                                                                 |     | 
| ReceivingPlantPartnerFunction  | 入荷プラント取引先機能。入出荷伝票から参照して登録される。変更不可。                                                                                                                 |     | 
| ReceivingPlantBusinessPartner  | 入荷プラントビジネスパートナ。入出荷伝票から参照して登録される。変更不可。                                                                                                           |     | 
| ReceivingPlant                 | 入荷プラント。入出荷伝票から参照して登録される。変更不可。                                                                                                                           |     | 
| ReceivingPlantStorageLocation  | 入荷プラントの保管場所。入出荷伝票から参照して登録される。変更不可。                                                                                                                 |     | 
| InvoiceDocumentItemText        | 請求伝票明細テキスト。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                             |     | 
| ServicesRenderedDate           | サービス提供日付。                                                                                                                                                                   |     | 
| InvoiceQuantity                | 請求数量。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                         |     | 
| InvoiceQuantityUnit            | 請求数量単位。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                     |     | 
| InvoiceQuantityInBaseUnit      | 請求数量（基本単位）。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                             |     | 
| BaseUnit                       | 基本単位。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                         |     | 
| ActualGoodsIssueDate           | 実際在庫出庫日付。入出荷伝票明細から参照して登録される。変更不可。                                                                                                                   |     | 
| ActualGoodsIssueTime           | 実際在庫出庫時刻。入出荷伝票明細から参照して登録される。変更不可。                                                                                                                   |     | 
| ActualGoodsReceiptDate         | 実際在庫入庫日付。入出荷伝票明細から参照して登録される。変更不可。                                                                                                                   |     | 
| ActualGoodsReceiptTime         | 実際在庫入庫時刻。入出荷伝票明細から参照して登録される。変更不可。                                                                                                                   |     | 
| ItemGrossWeight                | 明細総重量。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                       |     | 
| ItemNetWeight                  | 明細正味重量。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                     |     | 
| ItemWeightUnit                 | 明細重量単位。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                     |     | 
| ItemVolume                     | 明細数量。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                         |     | 
| ItemVolumeUnit                 | 明細数量単位。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                     |     | 
| NetAmount                      | 正味金額。消費税を除いた正味金額がオーダー明細から参照して登録される。必要に応じて変更する。                                                                                         |     | 
| GrossAmount                    | 総額。消費税を含んだ総額が自動計算される。必要に応じて変更する。総額と正味金額との差額が2消費税額として正しいかどうかチェックされる。理論値と2円以上の差額がある場合はエラーとなる。 |     | 
| TaxAmount                      | 消費税額。自動計算される。必要に応じて変更する。理論値と2円以上の差額がある場合はエラーとなる。                                                                                      |     | 
| GoodsIssueOrReceiptSlipNumber  | 納品書または受領書番号。入出庫伝票ヘッダから参照して登録される。変更不可。                                                                                                           |     | 
| TransactionCurrency            | 取引通貨。オーダー明細から参照して登録される。変更不可。                                                                                                                             |     | 
| PricingDate                    | 価格設定日付。オーダー明細から参照して登録される。変更不可。                                                                                                                         |     | 
| ProductTaxClassification       | 品目税分類。オーダー明細から参照して登録される。変更不可。                                                                                                                           |     | 
| Project                        | プロジェクト。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                     |     | 
| OrderID                        | オーダー番号。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                     |     | 
| OrderItem                      | オーダー明細番号。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                 |     | 
| OrderType                      | オーダータイプ。入出荷伝票明細またはオーダー明細から参照して登録される。                                                                                                             |     | 
| ContractType                   | 契約タイプ。オーダーから参照して登録される。変更不可。 <br> MNTH:月額契約。 <br>  ANNL:年額契約。                |     |                                                                                                                                                                                 | 
| OrderVaridityStartDate         | オーダー契約開始日付。オーダーから参照して登録される。変更不可。                                                                                                                     |     | 
| OrderValidityEndDate           | オーダー契約終了日付。オーダーから参照して登録される。変更不可。                                                                                                                     |     | 
| InvoiceScheduleStartDate       | 請求計画開始日付。オーダーから参照して登録される。変更不可。                                                                                                                         |     | 
| InvoiceScheduleEndDate         | 請求計画終了日付。オーダーから参照して登録される。変更不可。                                                                                                                         |     | 
| DeliveryDocument               | 入出荷伝票番号。参照された入出荷伝票番号。                                                                                                                                           |     | 
| DeliveryDocumentItem           | 入出荷伝票明細番号。参照された入出荷伝票明細番号。                                                                                                                                   |     | 
| OriginDocument                 | 原始伝票。変更不可。                                                                                                                                                                 |     | 
| OriginDocumentItem             | 原始伝票明細。変更不可。                                                                                                                                                             |     | 
| ReferenceDocument              | 参照伝票。変更不可。                                                                                                                                                                 |     | 
| ReferenceDocumentItem          | 参照伝票明細。変更不可。                                                                                                                                                             |     | 
| ReferenceDocumentType          | 参照伝票のタイプ。変更不可。                                                                                                                                                         |     | 
| ExternalReferenceDocument      | 外部参照伝票。周辺業務システムの参照伝票番号を入力する。                                                                                                                             |     | 
| ExternalReferenceDocumentItem  | 外部参照伝票明細。周辺業務システムの参照伝票明細番号を入力する。                                                                                                                     |     | 
| OrderDistributionChannel       | オーダー流通チャネル。オーダーから参照して登録される。変更不可。                                                                                                                     |     | 
| TaxCode                        | 消費税コード。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                     |     | 
| TaxRate                        | 消費税率。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                         |     | 
| CountryOfOrigin                | 原産国。入出荷伝票明細またはオーダー明細から参照して登録される。変更不可。                                                                                                           |     | 


### API名:【Invoice Document Item Partner > A_ItemPartner】
### Accepter名:【ItemPartner】

### <項目>
以下の表には、data-platform-api-invoice-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                              | Description                                                                                | EC  | 
| ------------------------------------- | ------------------------------------------------------------------------------------------ | --- | 
| PartnerFunction                       | 取引先機能。入出荷伝票またはオーダー明細から参照される。変更不可。 <br> 参照される取引先機能: <br> BUYER:買い手 <br> SELLER:売り手 <br> CUSTOMER:受注先 <br> SUPPLIER:仕入先 <br> MANUFACTURER:製造者 <br> DELIVERFROM:入出荷元 <br> DELIVERTO:入出荷先 <br> LOGI:物流業者 <br> BILLTO:請求先 <br> BILLFROM:請求元 <br> PAYEE:支払人 <br> RECEIVER:受取人 <br> PSPROVIDER:支払決済サービスプロバイダ |      |                                                                                      | 
| BusinessPartner                       | ビジネスパートナコード。取引先機能に対応するビジネスパートナコードが設定される。変更不可。 |     | 

### API名:【Invoice Document Item Pricing Element > A_ItemPricingElement】
### Accepter名:【ItemPricingElement】

### <項目>
以下の表には、data-platform-api-invoice-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property                   | Description                                                                      | EC  | 
| -------------------------- | -------------------------------------------------------------------------------- | --- | 
| PricingProcedureStep       | 価格手続ステップ。オーダーの明細価格決定要素から参照される。変更不可。           |     | 
| PricingProcedureCounter    | 価格手続カウンタ。オーダーの明細価格決定要素から参照される。変更不可。           |     | 
| ConditionType              | 条件タイプ。オーダーの明細価格決定要素から参照される。変更不可。                 |     | 
| PricingDate                | 価格設定日付。オーダーの明細価格決定要素から参照される。変更不可。               |     | 
| ConditionRateValue         | 条件レート値。オーダーの明細価格決定要素から参照される。変更不可。               |     | 
| ConditionCurrency          | 条件通貨。オーダーの明細価格決定要素から参照される。変更不可。                   |     | 
| ConditionQuantity          | 条件数量。オーダーの明細価格決定要素から参照される。変更不可。                   |     | 
| ConditionQuantityUnit      | 条件数量単位。オーダーの明細価格決定要素から参照される。変更不可。               |     | 
| ConditionRecord            | 条件レコード。オーダーの明細価格決定要素から参照される。変更不可。               |     | 
| ConditionSequentialNumber  | 条件連続番号。オーダーの明細価格決定要素から参照される。変更不可。               |     | 
| TaxCode                    | 消費税コード。オーダーの明細価格決定要素から参照される。変更不可。               |     | 
| ConditionAmount            | 条件金額。オーダーの明細価格決定要素から参照される。変更不可。                   |     | 
| TransactionCurrency        | 取引通貨。オーダーの明細価格決定要素から参照される。変更不可。                   |     | 
| ConditionIsManuallyChanged | 条件のマニュアル変更の有無。オーダーの明細価格決定要素から参照される。変更不可。 |     | 

### API名:【Invoice Document Address> A_Address】
### Accepter名:【Address】

### <項目>
以下の表には、data-platform-api-invoice-document-creates-rmq-kube のデータの項目とその定義が記載されています。  
※ EC の列には、そのデータの存在性確認を行うと定めているかどうか（✔がついていれば、行うと定めている）が記載されています。

| Property    | Description                                                                                                                    | EC  | 
| ----------- | ------------------------------------------------------------------------------------------------------------------------------ | --- | 
| AddressID   | 住所ID。オーダーまたは入出荷伝票から設定される。請求伝票でマニュアルで住所を更新する場合、住所IDが新たに設定される。変更不可。 |     | 
| PostalCode  | 郵便番号。オーダーまたは入出荷伝票から設定される。必要に応じて変更する。                                                       | ✔  | 
| LocalRegion | ローカル地域。オーダーまたは入出荷伝票から設定される。必要に応じて変更する。                                                   | ✔  | 
| Country     | 国。オーダーまたは入出荷伝票から設定される。必要に応じて変更する。                                                             | ✔  | 
| District    | ディストリクト。オーダーまたは入出荷伝票から設定される。必要に応じて変更する。                                                 | ✔  | 
| StreetName  | 地名・番地。オーダーまたは入出荷伝票から設定される。必要に応じて変更する。                                                     |     | 
| CityName    | 市区町村。オーダーまたは入出荷伝票から設定される。必要に応じて変更する。                                                       |     | 
| Building    | 建物名。オーダーまたは入出荷伝票から設定される。必要に応じて変更する。                                                         |     | 
| Floor       | 階数。オーダーまたは入出荷伝票から設定される。必要に応じて変更する。                                                           |     | 
| Room        | 部屋番号。オーダーまたは入出荷伝票から設定される。必要に応じて変更する。                                                       |     | 
