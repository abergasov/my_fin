<?php

use Phinx\Migration\AbstractMigration;

class RefreshTokenTable extends AbstractMigration
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
        $this->execute('create table user_refresh_tokens
                        (
                        	urt_id int auto_increment,
                        	refresh_token char(50) null,
                        	valid_until int null,
                        	fingerprint int null,
                        	user_id int null,
                        	constraint user_refresh_tokens_pk
                        		primary key (urt_id)
                        );');
        $this->execute('create index user_refresh_tokens_user_id_index
                        	on user_refresh_tokens (user_id);');
    }
}
