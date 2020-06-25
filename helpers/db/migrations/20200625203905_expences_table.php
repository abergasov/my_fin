<?php

use Phinx\Migration\AbstractMigration;

class ExpencesTable extends AbstractMigration
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
        $this->execute('create table expenses
                        (
                        	e_id int auto_increment,
                        	user_id int null,
                        	category smallint not null,
                        	amount float null,
                        	commentary varchar(255) null,
                        	type bit null,
                        	constraint expenses_pk
                        		primary key (e_id)
                        );');
        $this->execute('create index expenses_category_index on expenses (category);');
        $this->execute('create index expenses_type_index on expenses (type);');
        $this->execute('create index expenses_user_id_index on expenses (user_id);');
    }
}
