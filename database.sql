CREATE DATABASE scraper;

CREATE TABLE scrap (
    package_name        TEXT            PRIMARY KEY,
    app_name            TEXT,
    urlid               VARCHAR,
    company             VARCHAR,
    contain_ads         BOOLEAN,
    purchase_ads        BOOLEAN,
    image_url           VARCHAR,
    rating              VARCHAR,
    rated_people        VARCHAR,
    updated             VARCHAR,
    size                VARCHAR,
    installs            VARCHAR,
    current_version     VARCHAR,
    android_version     VARCHAR,
    content_rating      VARCHAR,
    interactive_ele     VARCHAR,
    in_app_products     VARCHAR,
    offered_by          VARCHAR,
    developer           VARCHAR
    );


-- CREATE TABLE scrap (package_name TEXT PRIMARY KEY,app_name TEXT NOT NULL,urlid VARCHAR NOT NULL,company VARCHAR NOT NULL,contain_ads BOOLEAN,purchase_ads BOOLEAN,image_url VARCHAR,rating VARCHAR,rated_people VARCHAR,updated VARCHAR,size VARCHAR,installs VARCHAR,current_version VARCHAR,android_version VARCHAR,content_rating VARCHAR,interactive_ele VARCHAR,in_app_products VARCHAR,offered_by VARCHAR,developer VARCHAR;


CREATE TABLE scrap_logs (
    package_name TEXT,
    field TEXT,
    old TEXT,
    new TEXT,
    updated DATE,
    UNIQUE(package_name,field,updated)
    );

create or replace function log_change()
returns trigger as
$BODY$
begin
if NEW.app_name<>OLD.app_name then
insert into scrap_logs values(old.package_name , 'app_name' , old.app_name , new.app_name, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.urlid<>OLD.urlid then
insert into scrap_logs values(old.package_name , 'urlid' , old.urlid , new.urlid, CURRENT_DATE)ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.company<>OLD.company then
insert into scrap_logs values(old.package_name , 'company' , old.company , new.company, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.contain_ads<>OLD.contain_ads then
insert into scrap_logs values(old.package_name , 'contain_ads' , old.contain_ads , new.contain_ads, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.purchase_ads<>OLD.purchase_ads then
insert into scrap_logs values(old.package_name , 'purchase_ads' , old.purchase_ads , new.purchase_ads, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.image_url<>OLD.image_url then
insert into scrap_logs values(old.package_name , 'image_url' , old.image_url , new.image_url, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.rating<>OLD.rating then
insert into scrap_logs values(old.package_name , 'rating' , old.rating , new.rating, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.rated_people<>OLD.rated_people then
insert into scrap_logs values(old.package_name , 'rated_people' , old.rated_people , new.rated_people, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.updated<>OLD.updated then
insert into scrap_logs values(old.package_name , 'updated' , old.updated , new.updated, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.size<>OLD.size then
insert into scrap_logs values(old.package_name , 'size' , old.size , new.size, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.installs<>OLD.installs then
insert into scrap_logs values(old.package_name , 'installs' , old.installs , new.installs, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.current_version<>OLD.current_version then
insert into scrap_logs values(old.package_name , 'current_version' , old.current_version , new.current_version, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.android_version<>OLD.android_version then
insert into scrap_logs values(old.package_name , 'android_version' , old.android_version , new.android_version, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.content_rating<>OLD.content_rating then
insert into scrap_logs values(old.package_name , 'content_rating' , old.content_rating , new.content_rating, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.interactive_ele<>OLD.interactive_ele then
insert into scrap_logs values(old.package_name , 'interactive_ele' , old.interactive_ele , new.interactive_ele, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.in_app_products<>OLD.in_app_products then
insert into scrap_logs values(old.package_name , 'in_app_products' , old.in_app_products , new.in_app_products, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.offered_by<>OLD.offered_by then
insert into scrap_logs values(old.package_name , 'offered_by' , old.offered_by , new.offered_by, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
if NEW.developer<>OLD.developer then
insert into scrap_logs values(old.package_name , 'developer' , old.developer , new.developer, CURRENT_DATE) ON CONFLICT (package_name,field,updated) DO UPDATE set old= EXCLUDED.old,new= EXCLUDED.new;
end if;
return new;
end;
$BODY$
language plpgsql;

create trigger log_changer
before update
on scrap
for each row
execute procedure log_change();
