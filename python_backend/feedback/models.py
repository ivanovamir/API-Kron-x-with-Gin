from django.db import models

# Create your models here.

class Supports(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.CharField(max_length=500, blank=True, null=True, verbose_name='Имя')
    body = models.TextField(blank=True, null=True, verbose_name='Запрос')

    class Meta:
        managed = False
        db_table = 'supports'
        verbose_name = 'Запрос на поддержку'
        verbose_name_plural = 'Запросы на поддержку'

    def __str__(self):
        return f'Запрос № - {self.id} | ФИО - {self.name}'


class FeedbackSelections(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.CharField(max_length=500, blank=True, null=True, verbose_name='Имя')
    phone = models.CharField(max_length=500, blank=True, null=True, verbose_name='Телефон')
    email = models.CharField(max_length=500, blank=True, null=True, verbose_name='Email')
    comment = models.TextField(blank=True, null=True, verbose_name='Запрос')

    class Meta:
        managed = False
        db_table = 'feedback_selections'
        verbose_name = 'Заявка на поиск деталей'
        verbose_name_plural = 'Заявки на поиск деталей'

    def __str__(self):
        return f'ФИО - {self.name}'


class FeedbackCalls(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.CharField(max_length=500, blank=True, null=True, verbose_name='Имя')
    phone = models.CharField(max_length=500, blank=True, null=True, verbose_name='Телефон')

    class Meta:
        managed = False
        db_table = 'feedback_calls'
        verbose_name = 'Обратный звонок'
        verbose_name_plural = 'Обратные звонки'

    def __str__(self):
        return f'ФИО - {self.name} | Телефон - {self.phone}'
