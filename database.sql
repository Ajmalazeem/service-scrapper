CREATE DATABASE scraper

ALTER TABLE scrap
CREATE TABLE scrap (
    id                  SERIAL          PRIMARY KEY,
    app_name            TEXT            NOT NULL,
    package_name        VARCHAR,
    urlid               VARCHAR         NOT NULL,
    company             VARCHAR         NOT NULL,
    contain_ads         BOOLEAN,
    purchase_ads        BOOLEAN,
    image_url           VARCHAR,
    rating              VARCHAR,
    rated_people        VARCHAR,
    updated             VARCHAR,
    size                VARCHAR,
    installs            VARCHAR,
    current_version     VARCHAR,
    android_version    VARCHAR,
    content_rating      VARCHAR,
    interactive_ele     VARCHAR,
    in_app_products     VARCHAR,
    offered_by          VARCHAR,
    developer           VARCHAR,
    UNIQUE(package_name)
    );



CREATE TABLE scrap (id                  SERIAL          PRIMARY KEY,app_name            TEXT            NOT NULL,package_name        VARCHAR,urlid               VARCHAR         NOT NULL,company             VARCHAR         NOT NULL,contain_ads         BOOLEAN,purchase_ads        BOOLEAN,image_url           VARCHAR,rating              VARCHAR,rated_people        VARCHAR,updated             VARCHAR,size                VARCHAR,installs            VARCHAR,current_version     VARCHAR,android_version    VARCHAR,content_rating      VARCHAR,interactive_ele     VARCHAR,in_app_products     VARCHAR,offered_by          VARCHAR,developer           VARCHAR,UNIQUE(package_name));
