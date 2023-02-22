--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.3
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
SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;
SET session_replication_role = 'replica';

--
-- Data for Name: core_users; Type: TABLE DATA; Schema: public; Owner: eunimartuser
--

INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (11, 'Super User', 'Super', 'User', 'super_user@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[1]', 1, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (12, 'Company Admin', 'Company', 'User', 'comapany_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[2]', NULL, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (13, 'MDM Admin', 'Mdm', 'Admin', 'mdm_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[3]', NULL, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (14, 'Order Admin', 'Order', 'Admin', 'order_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[4]', NULL, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (15, 'Inventory Admin', 'Inventory', 'Admin', 'inventory_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[5]', NULL, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (16, 'Invnetory Task Admin', 'Invnetory Task', 'Admin', 'inventory_task_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[6]', NULL, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (17, 'Accouting Admin', 'Accounting', 'Admin', 'accounting_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[7]', NULL, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (18, 'Returns Admin', 'LReturns', 'Admin', 'returns_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[8]', NULL, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (19, 'Shipping Admin', 'Shipping', 'Admin', 'shipping_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[9]', NULL, '{}', NULL, true, true, NULL, NULL, NULL, NULL,  '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (20, 'OmniChannel Admin', 'OmniChannel', 'Admin', 'omni_channel_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[10]', NULL, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 634,"name": "SELLER"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (22, 'BAP Admin', 'BAP', 'Admin', 'bap_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{"otp_token":"666666"}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[12]', 1, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 567,"name": "BAP_ADMIN"}]');
INSERT INTO public.core_users (id, name, first_name, last_name, username, email, work_email, mobile_number, password, login_type, auth, fa_conf, device_ids, preferences, tutorials, gamification, total_points, menu_hierarchy, constraints, schedulers, queue_services, access_ids, company_id, profile, plt_point_id, is_enabled, is_active, updated_date, created_date, created_by, updated_by, user_types) VALUES (23, 'BPP Admin', 'BPP', 'Admin', 'bpp_admin@eunimart.com', NULL, NULL, NULL, '', 0, '{"otp_token":"666666"}', '[]', '[]', '[]', '[]', '[]', 0, '{"otp_token": "666666"}', '[]', '[]', '[]', '[13]', 1, '{}', NULL, true, true, NULL, NULL, NULL, NULL, '[{"id": 568,"name": "BPP_ADMIN"}]');
SET session_replication_role = 'origin';

--
-- Name: core_users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: eunimartuser
--

select setval( pg_get_serial_sequence('public.core_users', 'id'), (select max(id) from public.core_users));

--
-- PostgreSQL database dump complete
--
