package eda

import (
	"os"
)

/*
 Copyright (C) 2022 Eunimart Omnichannel Pvt Ltd. (www.eunimart.com)
 All rights reserved.
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Lesser General Public License v3.0 as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Lesser General Public License v3.0 for more details.
 You should have received a copy of the GNU Lesser General Public License v3.0
 along with this program.  If not, see <https://www.gnu.org/licenses/lgpl-3.0.html/>.
*/

type TopicName string
type GroupName string

const (
	CREATE_CUSTOMER                          TopicName = "public.payments.customers.create"
	UPDATE_CUSTOMER                          TopicName = "public.payments.customers.update"
	CREATE_CREDITNOTE                        TopicName = "public.accounting.creditnote.create"
	UPDATE_CREDITNOTE                        TopicName = "public.accounting.creditnote.update"
	CREATE_DEBITNOTE                         TopicName = "public.accounting.debitnote.create"
	UPDATE_DEBITNOTE                         TopicName = "public.accounting.debitnote.update"
	CREATE_PURCHASE_INVOICE                  TopicName = "public.accounting.purchase_invoice.create"
	UPDATE_PURCHASE_INVOICE                  TopicName = "public.accounting.purchase_invoice.update"
	USER_LOGIN                               TopicName = "public.auth.user_login"
	VERIFY_OTP                               TopicName = "public.auth.verify_otp"
	CREATE_PRICING                           TopicName = "public.mdm.pricing.create"
	UPDATE_PRICING                           TopicName = "public.mdm.pricing.update"
	CREATE_TRANSACTION                       TopicName = "public.payments.transactions.create"
	UPDATE_TRANSACTION                       TopicName = "public.payments.transactions.update"
	CREATE_WALLET                            TopicName = "public.payments.wallets.create"
	UPDATE_WALLET                            TopicName = "public.payments.wallets.update"
	CREATE_PRODUCT_BRAND                     TopicName = "public.mdm.brand.create"
	UPDATE_PRODUCT_BRAND                     TopicName = "public.mdm.brand.update"
	CREATE_PRODUCT_VARIANT                   TopicName = "public.mdm.variant.create"
	UPDATE_PRODUCT_VARIANT                   TopicName = "public.mdm.variant.update"
	CREATE_CURRENCY_EXCHANGE                 TopicName = "public.accounting.currency_exchange.create"
	UPDATE_CURRENCY_EXCHANGE                 TopicName = "public.accounting.currency_exchange.update"
	CREATE_PRODUCT                           TopicName = "public.mdm.products.create"
	UPDATE_PRODUCT                           TopicName = "public.mdm.products.update"
	CREATE_CONTACTS                          TopicName = "public.mdm.contacts.create"
	UPDATE_CONTACTS                          TopicName = "public.mdm.contacts.update"
	CREATE_LOCATION                          TopicName = "public.mdm.locations.create"
	UPDATE_LOCATION                          TopicName = "public.mdm.locations.update"
	CREATE_INVENTORY                         TopicName = "public.mdm.inventory.create"
	UPDATE_INVENTORY                         TopicName = "public.mdm.inventory.update"
	UPDATE_COMPANY                           TopicName = "public.core.company.update"
	CREATE_SALES_ORDER                       TopicName = "public.orders.sales_order.create"
	UPDATE_SALES_ORDER                       TopicName = "public.orders.sales_order.update"
	CREATE_SHIPPING_ORDER                    TopicName = "public.shipping.shipping_orders.create"
	UPDATE_SHIPPING_ORDER                    TopicName = "public.shipping.shipping_orders.update"
	CREATE_SHIPPING_ORDER_NDR                TopicName = "public.shipping.shipping_orders_ndr.create"
	UPDATE_SHIPPING_ORDER_NDR                TopicName = "public.shipping.shipping_orders_ndr.update"
	CREATE_SHIPPING_ORDER_RTO                TopicName = "public.shipping.shipping_orders_rto.create"
	UPDATE_SHIPPING_ORDER_RTO                TopicName = "public.shipping.shipping_orders_rto.update"
	CREATE_SHIPPING_ORDER_WD                 TopicName = "public.shipping.shipping_orders_wd.create"
	UPDATE_SHIPPING_ORDER_WD                 TopicName = "public.shipping.shipping_orders_wd.update"
	CREATE_SHIPPING_PARTNER                  TopicName = "public.shipping.shipping_partner.create"
	UPDATE_SHIPPING_PARTNER                  TopicName = "public.shipping.shipping_partner.update"
	CREATE_PURCHASE_RETURNS                  TopicName = "public.returns.purchase_returns.create"
	UPDATE_PURCHASE_RETURNS                  TopicName = "public.returns.purchase_returns.update"
	CREATE_SALES_RETURNS                     TopicName = "public.returns.sales_returns.create"
	UPDATE_SALES_RETURNS                     TopicName = "public.returns.sales_returns.update"
	CREATE_INVENTORY_ADJUSTMENTS             TopicName = "public.inventory_orders.inventory_adjustments.create"
	UPDATE_INVENTORY_ADJUSTMENTS             TopicName = "public.inventory_orders.inventory_adjustments.update"
	CREATE_ASN                               TopicName = "public.inventory_orders.asn.create"
	UPDATE_ASN                               TopicName = "public.inventory_orders.asn.update"
	CREATE_GRN                               TopicName = "public.inventory_orders.grn.create"
	UPDATE_GRN                               TopicName = "public.inventory_orders.grn.update"
	CREATE_UOM                               TopicName = "public.mdm.uom.create"
	UPDATE_UOM                               TopicName = "public.mdm.uom.update"
	CREATE_UOM_CLASS                         TopicName = "public.mdm.uom_class.create"
	UPDATE_UOM_CLASS                         TopicName = "public.mdm.uom_class.update"
	BOSON_MAPPER_CONVERTOR                   TopicName = "public.ipaas_core.boson_convertor"
	BOSON_MAPPER_CONVERTOR_ACK               TopicName = "public.tech_ack"
	CREATE_PICK_LIST                         TopicName = "public.inventory_tasks.pick_list.create"
	UPDATE_PICK_LIST                         TopicName = "public.inventory_tasks.pick_list.update"
	CREATE_CYCLE_COUNT                       TopicName = "public.inventory_tasks.cycle_count.create"
	UPDATE_CYCLE_COUNT                       TopicName = "public.inventory_tasks.cycle_count.update"
	CREATE_SALES_INVOICE                     TopicName = "public.accounting.sales_invoice.create"
	UPDATE_SALES_INVOICE                     TopicName = "public.accounting.sales_invoice.update"
	CREATE_INTERNAL_TRANSFERS                TopicName = "public.orders.internal_transfers.create"
	UPDATE_INTERNAL_TRANSFERS                TopicName = "public.orders.internal_transfers.update"
	CREATE_SCRAP_ORDERS                      TopicName = "public.orders.scrap_orders.create"
	UPDATE_SCRAP_ORDERS                      TopicName = "public.orders.scrap_orders.update"
	CREATE_PURCHASE_ORDERS                   TopicName = "public.orders.purchase_orders.create"
	UPDATE_PURCHASE_ORDERS                   TopicName = "public.orders.purchase_orders.update"
	CREATE_DELIVERY_ORDERS                   TopicName = "public.orders.delivery_orders.create"
	UPDATE_DELIVERY_ORDERS                   TopicName = "public.orders.delivery_orders.update"
	CREATE_PRODUCT_BUNDLE                    TopicName = "public.mdm.products_bundle.create"
	UPDATE_PRODUCT_BUNDLE                    TopicName = "public.mdm.products_bundle.update"
	CREATE_PRODUCT_CATEGORY                  TopicName = "public.mdm.products_category.create"
	UPDATE_PRODUCT_CATEGORY                  TopicName = "public.mdm.products_category.update"
	CREATE_PAYMENT_TERMS_AND_RECORD_PAYMENTS TopicName = "public.accounting.payment_terms_and_record_payment.create"
	UPDATE_PAYMENT_TERMS_AND_RECORD_PAYMENTS TopicName = "public.accounting.payment_terms_and_record_payment.update"
	CREATE_VENDOR                            TopicName = "public.mdm.vendor.create"
	UPDATE_VENDOR                            TopicName = "public.mdm.vendor.update"
	UPSERT_DECENTRALIZED_INVENTORY           TopicName = "public.mdm.decentralized_inventory_upsert"
	UPSERT_SALES_PRICE_LIST                  TopicName = "public.mdm.pricing.sales_price_list_upsert"
	UPSERT_PRODUCT_TEMPLATE                  TopicName = "public.mdm.products.template_upsert"
	UPSERT_PRODUCT_VARIANT                   TopicName = "public.mdm.products.variant_upsert"
	UPSERT_SALES_ORDER                       TopicName = "public.orders.sales_order.upsert"
	UPSERT_VENDOR                            TopicName = "public.mdm.vendor.upsert"
	UPSERT_CONTACT                           TopicName = "public.mdm.contact.upsert"
	CREATE_VENDOR_ACK                        TopicName = "public.tech_ack"
	UPDATE_VENDOR_ACK                        TopicName = "public.tech_ack"
	CREATE_INVENTORY_ACK                     TopicName = "public.tech_ack"
	UPDATE_INVENTORY_ACK                     TopicName = "public.tech_ack"
	CREATE_CONTACTS_ACK                      TopicName = "public.tech_ack"
	UPDATE_CONTACTS_ACK                      TopicName = "public.tech_ack"
	BECKN_PRODUCTS_SYNC                      TopicName = "public.ondc-beckn.products.sync"
	BECKN_PRODUCTS_SYNC_ACK                  TopicName = "public.ondc-beckn.products.sync_ack"
	BECKN_PROVIDER_SYNC                      TopicName = "public.ondc-beckn.core.company.sync"
	BECKN_PROVIDER_SYNC_ACK                  TopicName = "public.ondc-beckn.core.company.sync_ack"

	UPSERT_CONTACTS_ACK                          TopicName = "public.tech_ack"
	UPSERT_VENDORS_ACK                           TopicName = "public.tech_ack"
	CREATE_DELIVERY_ORDERS_ACK                   TopicName = "public.tech_ack"
	UPDATE_DELIVERY_ORDERS_ACK                   TopicName = "public.tech_ack"
	CREATE_INTERNAL_TRANSFERS_ACK                TopicName = "public.tech_ack"
	UPDATE_INTERNAL_TRANSFERS_ACK                TopicName = "public.tech_ack"
	CREATE_PURCHASE_ORDERS_ACK                   TopicName = "public.teck_ack"
	UPDATE_PURCHASE_ORDERS_ACK                   TopicName = "public.teck_ack"
	CREATE_SCRAP_ORDERS_ACK                      TopicName = "public.tech_ack"
	UPDATE_SCRAP_ORDERS_ACK                      TopicName = "public.tech_ack"
	CREATE_SALES_ORDER_ACK                       TopicName = "public.tech_ack"
	UPDATE_SALES_ORDER_ACK                       TopicName = "public.tech_ack"
	CREATE_PRODUCT_ACK                           TopicName = "public.tech_ack"
	UPDATE_PRODUCT_ACK                           TopicName = "public.tech_ack"
	UPDATE_PRODUCT_BRAND_ACK                     TopicName = "public.tech_ack"
	CREATE_PRODUCT_BRAND_ACK                     TopicName = "public.tech_ack"
	CREATE_PRODUCT_VARIANT_ACK                   TopicName = "public.tech_ack"
	UPDATE_PRODUCT_VARIANT_ACK                   TopicName = "public.tech_ack"
	CREATE_PRODUCT_CATEGORY_ACK                  TopicName = "public.tech_ack"
	UPDATE_PRODUCT_CATEGORY_ACK                  TopicName = "public.tech_ack"
	CREATE_PRODUCT_BUNDLE_ACK                    TopicName = "public.tech_ack"
	UPDATE_PRODUCT_BUNDLE_ACK                    TopicName = "public.tech_ack"
	UPSERT_DECENTRALIZED_INVENTORY_ACK           TopicName = "public.tech_ack"
	UPSERT_SALES_ORDER_ACK                       TopicName = "public.tech_ack"
	UPSERT_PRODUCT_VARIANT_ACK                   TopicName = "public.tech_ack"
	UPSERT_PRODUCT_TEMPLATE_ACK                  TopicName = "public.tech_ack"
	CREATE_UOM_CLASS_ACK                         TopicName = "public.tech_ack"
	UPDATE_UOM_CLASS_ACK                         TopicName = "public.tech_ack"
	CREATE_UOM_ACK                               TopicName = "public.tech_ack"
	UPDATE_UOM_ACK                               TopicName = "public.tech_ack"
	CREATE_ASN_ACK                               TopicName = "public.tech_ack"
	UPDATE_ASN_ACK                               TopicName = "public.tech_ack"
	CREATE_GRN_ACK                               TopicName = "public.tech_ack"
	UPDATE_GRN_ACK                               TopicName = "public.tech_ack"
	UPDATE_INVENTORY_ADJUSTMENTS_ACK             TopicName = "public.tech_ack"
	CREATE_INVENTORY_ADJUSTMENTS_ACK             TopicName = "public.tech_ack"
	CREATE_CREDITNOTE_ACK                        TopicName = "public.tech_ack"
	UPDATE_CREDITNOTE_ACK                        TopicName = "public.tech_ack"
	CREATE_DEBITNOTE_ACK                         TopicName = "public.tech_ack"
	UPDATE_DEBITNOTE_ACK                         TopicName = "public.tech_ack"
	CREATE_CURRENCY_EXCHANGE_ACK                 TopicName = "public.tech_ack"
	UPDATE_CURRENCY_EXCHANGE_ACK                 TopicName = "public.tech_ack"
	CREATE_PURCHASE_RETURNS_ACK                  TopicName = "public.tech_ack"
	UPDATE_PURCHASE_RETURNS_ACK                  TopicName = "public.tech_ack"
	CREATE_SALES_RETURNS_ACK                     TopicName = "public.tech_ack"
	UPDATE_SALES_RETURNS_ACK                     TopicName = "public.tech_ack"
	CREATE_CUSTOMER_ACK                          TopicName = "public.tech_ack"
	UPDATE_CUSTOMER_ACK                          TopicName = "public.tech_ack"
	CREATE_TRANSACTION_ACK                       TopicName = "public.tech_ack"
	UPDATE_TRANSACTION_ACK                       TopicName = "public.tech_ack"
	CREATE_WALLET_ACK                            TopicName = "public.tech_ack"
	UPDATE_WALLET_ACK                            TopicName = "public.tech_ack"
	CREATE_SALES_INVOICE_ACK                     TopicName = "public.tech_ack"
	UPDATE_SALES_INVOICE_ACK                     TopicName = "public.tech_ack"
	CREATE_PURCHASE_INVOICE_ACK                  TopicName = "public.tech_ack"
	UPDATE_PURCHASE_INVOICE_ACK                  TopicName = "public.tech_ack"
	CREATE_PAYMENT_TERMS_AND_RECORD_PAYMENTS_ACK TopicName = "public.tech_ack"
	UPDATE_PAYMENT_TERMS_AND_RECORD_PAYMENTS_ACK TopicName = "public.tech_ack"
	CREATE_PICK_LIST_ACK                         TopicName = "public.tech_ack"
	UPDATE_PICK_LIST_ACK                         TopicName = "public.tech_ack"
	CREATE_CYCLE_COUNT_ACK                       TopicName = "public.tech_ack"
	UPDATE_CYCLE_COUNT_ACK                       TopicName = "public.tech_ack"
	USER_LOGIN_ACK                               TopicName = "public.tech_ack"
	VERIFY_OTP_ACK                               TopicName = "public.tech_ack"
	UPDATE_COMPANY_ACK                           TopicName = "public.tech_ack"
	CREATE_LOCATION_ACK                          TopicName = "public.tech_ack"
	UPDATE_LOCATION_ACK                          TopicName = "public.tech_ack"
	CREATE_PRICING_ACK                           TopicName = "public.tech_ack"
	UPDATE_PRICING_ACK                           TopicName = "public.tech_ack"
	UPSERT_SALES_PRICE_LIST_ACK                  TopicName = "public.tech_ack"
	CREATE_SHIPPING_ORDER_ACK                    TopicName = "public.tech_ack"
	UPDATE_SHIPPING_ORDER_ACK                    TopicName = "public.tech_ack"
	CREATE_SHIPPING_ORDER_NDR_ACK                TopicName = "public.tech_ack"
	UPDATE_SHIPPING_ORDER_NDR_ACK                TopicName = "public.tech_ack"
	CREATE_SHIPPING_ORDER_RTO_ACK                TopicName = "public.tech_ack"
	UPDATE_SHIPPING_ORDER_RTO_ACK                TopicName = "public.tech_ack"
	CREATE_SHIPPING_ORDER_WD_ACK                 TopicName = "public.tech_ack"
	UPDATE_SHIPPING_ORDER_WD_ACK                 TopicName = "public.tech_ack"
	CREATE_SHIPPING_PARTNER_ACK                  TopicName = "public.tech_ack"
	UPDATE_SHIPPING_PARTNER_ACK                  TopicName = "public.tech_ack"
	CREATE_CENTRALIZED_INVENTORY                 TopicName = "public.mdm.centralized_inventory.create"
	UPDATE_CENTRALIZED_INVENTORY                 TopicName = "public.mdm.centralized_inventory.update"
	CREATE_CENTRALIZED_INVENTORY_ACK             TopicName = "public.tech_ack"
	UPDATE_CENTRALIZED_INVENTORY_ACK             TopicName = "public.tech_ack"
	CREATE_RATING                                TopicName = "public.rating.rating.create"
	CREATE_RATING_ACK                            TopicName = "public.tech_ack"
	UPDATE_RATING                                TopicName = "public.rating.rating.update"
	UPDATE_RATING_ACK                            TopicName = "public.tech_ack"
	CREATE_FEEDBACKFORM                          TopicName = "public.rating.rating.create_feedbackform"
	CREATE_FEEDBACKFORM_ACK                      TopicName = "public.tech_ack"
	CREATE_FEEDBACK                              TopicName = "public.rating.rating.create_feedback"
	CREATE_FEEDBACK_ACK                          TopicName = "public.tech_ack"
	UPDATE_FEEDBACK                              TopicName = "public.rating.rating.update_feedback"
	UPDATE_FEEDBACK_ACK                          TopicName = "public.tech_ack"
	CREATE_REPORT                                TopicName = "public.rating.rating.create_report"
	CREATE_REPORT_ACK                            TopicName = "public.tech_ack"
	UPDATE_REPORT                                TopicName = "public.rating.rating.update_report"
	UPDATE_REPORT_ACK                            TopicName = "public.tech_ack"
	CREATE_OFFER                                 TopicName = "public.offers.offers.create"
	CREATE_OFFER_ACK                             TopicName = "public.tech_ack"
	UPDATE_OFFER                                 TopicName = "public.offers.offers.update"
	UPDATE_OFFER_ACK                             TopicName = "public.tech_ack"

	CANCEL_SALES_ORDER             TopicName = "public.orders.sales_order.cancel"
	ONDC_UPDATE_SALES_ORDER_STATUS TopicName = "public.orders.sales_order.ondc_update_status"
	ONDC_CREATE_ORDER              TopicName = "public.ondc.orders.create"
	UPDATE_SALES_ORDER_STATUS 	   TopicName = "public.ondc.orders.update_status"
	GET_SALES_ORDER                TopicName = "public.orders.sales_order.status"
	SEND_SALES_ORDER               TopicName = "public.ondc.orders.status"
	CANCEL_SALES_ORDER_RESPONSE    TopicName = "public.ondc.orders.cancel"
)

const (
	BASE_GROUP       GroupName = "tech.consumer.group"
	MDM              GroupName = "tech.consumer.group"
	ORDERS           GroupName = "tech.consumer.group"
	INVENTORY        GroupName = "tech.consumer.group"
	CORE             GroupName = "tech.consumer.group"
	SHIPPING_ORDER   GroupName = "tech.consumer.group"
	RETURNS          GroupName = "tech.consumer.group"
	INVENTORY_ORDERS GroupName = "tech.consumer.group"
	PAYMENTS         GroupName = "tech.consumer.group"
	BOSON_CONVERTOR  GroupName = "tech.consumer.group"
	INVENTORY_TASKS  GroupName = "tech.consumer.group"
	ACCOUNTING       GroupName = "tech.consumer.group"
	BECKN            GroupName = "tech.consumer.group"
	RATING           GroupName = "tech.consumer.group"
	OFFERS           GroupName = "tech.consumer.group"
)

// Returns the Topic Name with the topic prefix appended to it
func GetTechTopicName(topic TopicName) TopicName {
	return TopicName(os.Getenv("TOPIC_PREFIX")+".") + topic
}

func GetOndcTopicName(topic TopicName) TopicName {
	return TopicName(os.Getenv("ONDC_PREFIX")+".") + topic
}

type ConsumerResponse struct {
	MetaData     interface{} `json:"meta_data"`
	ErrorMessage error       `json:"error_message"`
	Response     interface{} `json:"response"`
}