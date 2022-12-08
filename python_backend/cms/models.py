from django.db import models


class Emails(models.Model):
    id = models.BigAutoField(primary_key=True)
    logo_image = models.ImageField(upload_to="data/photos/email_data", blank=True, null=True,
                                   verbose_name='Лого компании')
    cart_image = models.ImageField(upload_to="data/photos/email_data", blank=True, null=True,
                                   verbose_name='Картинка "Корзины"')
    check_image = models.ImageField(upload_to="data/photos/email_data", blank=True, null=True,
                                    verbose_name='Картинка "Галочки"')
    address = models.TextField(blank=True, null=True, verbose_name='Адрес')

    class Meta:
        managed = False
        db_table = 'emails'
        verbose_name = 'Данные для Email уведомлений о заказах'
        verbose_name_plural = 'Данные для Email уведомлений о заказах'

    def __str__(self):
        return f'Данные Email - {self.id}'


class Sliders(models.Model):
    id = models.BigAutoField(primary_key=True)
    main_text = models.CharField(max_length=300, blank=True, null=True, verbose_name='Главный текст')
    upper_text = models.CharField(max_length=300, blank=True, null=True, verbose_name='Текст выше')
    down_text = models.TextField(blank=True, null=True, verbose_name='Текст ниже')
    image = models.ImageField(upload_to="data/photos/sliders", null=True, verbose_name='Фотография')
    link = models.CharField(max_length=500, blank=True, null=True, verbose_name='Ссылка')

    class Meta:
        managed = False
        db_table = 'sliders'
        verbose_name = 'Слайдер'
        verbose_name_plural = 'Слайдеры'
        ordering = ['id']

    def __str__(self):
        return f'{self.main_text}'


class Features(models.Model):
    id = models.BigAutoField(primary_key=True)
    header = models.CharField(max_length=300, blank=True, null=True, verbose_name='Заголовок')
    body = models.TextField(blank=True, null=True, verbose_name='Текст')
    icon = models.ImageField(blank=True, null=True, upload_to="data/photos/features", verbose_name='Иконка')

    class Meta:
        managed = False
        db_table = 'features'
        verbose_name = 'Достоинство'
        verbose_name_plural = 'Достоинства'

    def __str__(self):
        return f'{self.header}'


class Services(models.Model):
    id = models.BigAutoField(primary_key=True)
    title = models.CharField(max_length=500, blank=True, null=True, verbose_name='Название')
    body = models.TextField(blank=True, null=True, verbose_name='Содержание')
    image = models.ImageField(upload_to="data/photos/services", null=True, verbose_name='Фотография')

    class Meta:
        managed = False
        db_table = 'services'
        verbose_name = 'Услуга'
        verbose_name_plural = 'Услуги'

    def __str__(self):
        return f'Имя услуги - {self.title}'


class Metrics(models.Model):
    id = models.BigAutoField(primary_key=True)
    google_metric = models.CharField(max_length=500, blank=True, null=True, verbose_name='Google метрика')
    yandex_metric = models.CharField(max_length=500, blank=True, null=True, verbose_name='Yandex метрика')

    class Meta:
        managed = False
        db_table = 'metrics'
        verbose_name = 'Метрика'
        verbose_name_plural = 'Метрика'

    def __str__(self):
        return f'{self.id}'
