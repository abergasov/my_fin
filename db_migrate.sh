git fetch --all
git reset --hard origin/master


echo "migrate database"

cd helpers || exit
composer install
if php app.php migrate ; then
    echo "migrated"
else
    echo "migrated failed"
    exit 1
fi