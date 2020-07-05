<?php

use Phinx\Migration\AbstractMigration;

class UserAssets extends AbstractMigration
{
    /**
     * Change Method.
     *
     * Write your reversible migrations using this method.
     *
     * More information on writing migrations is available here:
     * http://docs.phinx.org/en/latest/migrations.html#the-abstractmigration-class
     *
     * The following commands can be used in this method and Phinx will
     * automatically reverse them when rolling back:
     *
     *    createTable
     *    renameTable
     *    addColumn
     *    addCustomColumn
     *    renameColumn
     *    addIndex
     *    addForeignKey
     *
     * Any other destructive changes will result in an error when trying to
     * rollback the migration.
     *
     * Remember to call "create()" or "update()" and NOT "save()" when working
     * with the Table class.
     */
    public function change()
    {
        $this->execute('create table `user_financial assets`
                        (
                        	ufa_id int auto_increment,
                        	user_id int null,
                        	asset_type char null,
                        	asset_name varchar(50) null,
                        	amount int null,
                        	constraint `user_financial assets_pk`
                        		primary key (ufa_id)
                        );    ');
        $this->execute('create index `user_financial assets_asset_type_index`
                                                	on `user_financial assets` (asset_type);');
        $this->execute('create index `user_financial assets_user_id_index`
                                                	on `user_financial assets` (user_id);');
    }
}
