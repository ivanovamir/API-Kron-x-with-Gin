# Generated by Django 4.0.5 on 2022-11-07 14:56

from django.db import migrations, models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
    ]

    operations = [
        migrations.CreateModel(
            name='FeedbackCalls',
            fields=[
                ('id', models.BigAutoField(primary_key=True, serialize=False)),
                ('name', models.CharField(blank=True, max_length=500, null=True, verbose_name='Имя')),
                ('phone', models.CharField(blank=True, max_length=500, null=True, verbose_name='Телефон')),
            ],
            options={
                'verbose_name': 'Обратный звонок',
                'verbose_name_plural': 'Обратные звонки',
                'db_table': 'feedback_calls',
                'managed': False,
            },
        ),
        migrations.CreateModel(
            name='FeedbackSelections',
            fields=[
                ('id', models.BigAutoField(primary_key=True, serialize=False)),
                ('name', models.CharField(blank=True, max_length=500, null=True, verbose_name='Имя')),
                ('phone', models.CharField(blank=True, max_length=500, null=True, verbose_name='Телефон')),
                ('email', models.CharField(blank=True, max_length=500, null=True, verbose_name='Email')),
                ('comment', models.TextField(blank=True, null=True, verbose_name='Запрос')),
            ],
            options={
                'verbose_name': 'Заявка на поиск деталей',
                'verbose_name_plural': 'Заявки на поиск деталей',
                'db_table': 'feedback_selections',
                'managed': False,
            },
        ),
        migrations.CreateModel(
            name='Supports',
            fields=[
                ('id', models.BigAutoField(primary_key=True, serialize=False)),
                ('name', models.CharField(blank=True, max_length=500, null=True, verbose_name='Имя')),
                ('body', models.TextField(blank=True, null=True, verbose_name='Запрос')),
            ],
            options={
                'verbose_name': 'Запрос на поддержку',
                'verbose_name_plural': 'Запросы на поддержку',
                'db_table': 'supports',
                'managed': False,
            },
        ),
    ]
