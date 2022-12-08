# Generated by Django 4.0.5 on 2022-11-07 14:56

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='Abouts',
            fields=[
                ('id', models.BigAutoField(primary_key=True, serialize=False)),
                ('main_image', models.ImageField(blank=True, null=True, upload_to='data/photos/about', verbose_name='Главная фотография')),
                ('secondary_image', models.ImageField(blank=True, null=True, upload_to='data/photos/about', verbose_name='Второстепенное фотография')),
                ('image1', models.ImageField(blank=True, null=True, upload_to='data/photos/about', verbose_name='Фото 1')),
                ('image2', models.ImageField(blank=True, null=True, upload_to='data/photos/about', verbose_name='Фото 2')),
                ('image3', models.ImageField(blank=True, null=True, upload_to='data/photos/about', verbose_name='Фото 3')),
                ('image4', models.ImageField(blank=True, null=True, upload_to='data/photos/about', verbose_name='Фото 4')),
                ('image5', models.ImageField(blank=True, null=True, upload_to='data/photos/about', verbose_name='Фото 5')),
                ('image6', models.ImageField(blank=True, null=True, upload_to='data/photos/about', verbose_name='Фото 6')),
                ('image7', models.ImageField(blank=True, null=True, upload_to='data/photos/about', verbose_name='Фото 7')),
            ],
            options={
                'verbose_name': 'О компании',
                'verbose_name_plural': 'О компании',
                'db_table': 'abouts',
                'managed': False,
            },
        ),
        migrations.CreateModel(
            name='HeadOffices',
            fields=[
                ('id', models.BigAutoField(primary_key=True, serialize=False)),
                ('title', models.TextField(blank=True, null=True, verbose_name='Название')),
                ('phone', models.TextField(blank=True, null=True, verbose_name='Телефон')),
                ('mail', models.TextField(blank=True, null=True, verbose_name='Почта')),
                ('schedule', models.TextField(blank=True, null=True, verbose_name='График')),
                ('address', models.TextField(blank=True, null=True, verbose_name='Адресс')),
            ],
            options={
                'verbose_name': 'Офис',
                'verbose_name_plural': 'Офисы',
                'db_table': 'head_offices',
                'managed': False,
            },
        ),
        migrations.CreateModel(
            name='Requisites',
            fields=[
                ('id', models.BigAutoField(primary_key=True, serialize=False)),
                ('company_name', models.CharField(blank=True, max_length=500, null=True, verbose_name='Имя компании')),
                ('address', models.TextField(blank=True, null=True, verbose_name='Адресс')),
                ('inn', models.CharField(blank=True, max_length=500, null=True, verbose_name='ИНН')),
                ('kpp', models.CharField(blank=True, max_length=500, null=True, verbose_name='КПП')),
                ('ogrn', models.CharField(blank=True, max_length=500, null=True, verbose_name='ОГРН')),
                ('rs', models.CharField(blank=True, max_length=500, null=True, verbose_name='Р/С')),
                ('bank', models.TextField(blank=True, null=True, verbose_name='Банк')),
                ('ks', models.CharField(blank=True, max_length=500, null=True, verbose_name='К/С')),
                ('bik', models.CharField(blank=True, max_length=500, null=True, verbose_name='БИК')),
            ],
            options={
                'verbose_name': 'Реквизит',
                'verbose_name_plural': 'Реквизиты',
                'db_table': 'requisites',
                'managed': False,
            },
        ),
    ]