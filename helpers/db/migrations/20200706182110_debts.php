<?php

use Phinx\Migration\AbstractMigration;

class Debts extends AbstractMigration
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
        $this->execute("create table debts
                        (
                        	d_id int auto_increment,
                        	user_id int not null,
                        	amount int null,
                        	until_date int null,
                        	active_debt enum('1', '0') default '1' null,
                        	constraint debts_pk
                        		primary key (d_id)
                        );");
        $this->execute('create index debts_active_debt_index
                        	on debts (active_debt);');
        $this->execute('create index debts_until_date_index
                        	on debts (until_date);');
        $this->execute('create index debts_user_id_index
                        	on debts (user_id);');
    }
}
