

---------MENUU

INSERT INTO m_menu_menuitems (menu_id, menu_key, menu_order, menu_url, menu_parent, menu_icon)
SELECT NEXT VALUE FOR seq_MENU, 'configuration_folder', 5, null, null, null
WHERE NOT EXISTS (SELECT menu_id FROM m_menu_menuitems WHERE menu_key = 'configuration_folder' );

INSERT INTO m_menu_menuitems_i18n (menui18n_id, menu_description, menu_name, menu_locale, menu_idmenu)
SELECT NEXT VALUE FOR seq_MENUI18N, 'Configurazione', 'Configurazione', (SELECT l.loca_id FROM m_loca_locales l WHERE l.loca_code = 'IT'),(SELECT menu_id FROM m_menu_menuitems WHERE menu_key = 'configuration_folder')
    WHERE NOT EXISTS (SELECT menui18n_id FROM m_menu_menuitems_i18n WHERE menu_idmenu = (SELECT menu_id FROM m_menu_menuitems WHERE menu_key = 'configuration_folder' )
  AND MENU_LOCALE = (SELECT l.loca_id FROM m_loca_locales l WHERE l.loca_code = 'IT') );
INSERT INTO m_menu_menuitems_i18n (menui18n_id, menu_description, menu_name, menu_locale, menu_idmenu)
SELECT NEXT VALUE FOR seq_MENUI18N, 'Configuration', 'Configuration', (SELECT l.loca_id FROM m_loca_locales l WHERE l.loca_code = 'EN'), (SELECT menu_id FROM m_menu_menuitems WHERE menu_key = 'configuration_folder')
WHERE NOT EXISTS (SELECT menui18n_id FROM m_menu_menuitems_i18n WHERE menu_idmenu = (SELECT menu_id FROM m_menu_menuitems WHERE menu_key = 'configuration_folder' )
  AND MENU_LOCALE = (SELECT l.loca_id FROM m_loca_locales l WHERE l.loca_code = 'EN'));


  

  EXEC P_SYNCRONIZE_SEQUENCE @SeqName = 'seq_cons', @ColId = 'cons_id',@TableName = 'm_cons_constants' ;
EXEC P_SYNCRONIZE_SEQUENCE @SeqName = 'seq_consi18n', @ColId = 'consi18n_id',@TableName = 'm_cons_constants_i18n' ;


EXEC sp_rename 'm_actn_action.actn_account_type' , 'actn_account_type_toremove','COLUMN';
