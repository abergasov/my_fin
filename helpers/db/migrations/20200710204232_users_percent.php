<?php

use Phinx\Migration\AbstractMigration;

class UsersPercent extends AbstractMigration
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
        //обязательные расходы
        $this->execute("alter table users add mandatory_percent tinyint default 30 not null;");
        //РНЖ
        $this->execute("alter table users add live_percent tinyint default 20 not null;");
        //Откладывание подушки
        $this->execute("alter table users add black_day_percent tinyint default 15 not null;");
        //Инвестирование
        $this->execute("alter table users add invest_percent tinyint default 15 not null;");
        //траты
        $this->execute("alter table users add spending_percent tinyint default 20 not null;");
    }
}
