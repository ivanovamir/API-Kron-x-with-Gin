from django.db import models

# Create your models here.

class HeadOffices(models.Model):
    id = models.BigAutoField(primary_key=True)
    title = models.TextField(blank=True, null=True, verbose_name='Название')
    phone = models.TextField(blank=True, null=True, verbose_name='Телефон')
    mail = models.TextField(blank=True, null=True, verbose_name='Почта')
    schedule = models.TextField(blank=True, null=True, verbose_name='График')
    address = models.TextField(blank=True, null=True, verbose_name='Адресс')

    class Meta:
        managed = False
        db_table = 'head_offices'
        verbose_name = 'Офис'
        verbose_name_plural = 'Офисы'

    def __str__(self):
        return f'Офис - {self.title}'


class Requisites(models.Model):
    id = models.BigAutoField(primary_key=True)
    company_name = models.CharField(max_length=500 ,blank=True, null=True, verbose_name='Имя компании')
    address = models.TextField(blank=True, null=True, verbose_name='Адресс')
    inn = models.CharField(max_length=500 ,blank=True, null=True, verbose_name='ИНН')
    kpp = models.CharField(max_length=500 , blank=True, null=True , verbose_name='КПП')
    ogrn = models.CharField(max_length=500 ,blank=True, null=True, verbose_name='ОГРН')
    rs = models.CharField(max_length=500 ,blank=True, null=True, verbose_name='Р/С')
    bank = models.TextField(blank=True, null=True, verbose_name='Банк')
    ks = models.CharField(max_length=500 ,blank=True, null=True, verbose_name='К/С')
    bik = models.CharField(max_length=500 ,blank=True, null=True, verbose_name='БИК')

    class Meta:
        managed = False
        db_table = 'requisites'
        verbose_name = 'Реквизит'
        verbose_name_plural = 'Реквизиты'

    def __str__(self):
        return f'Реквизит'


class Abouts(models.Model):
    id = models.BigAutoField(primary_key=True)
    main_image = models.ImageField(blank=True, null=True, upload_to="data/photos/about", verbose_name = 'Главная фотография')
    secondary_image = models.ImageField(blank=True, null=True, upload_to="data/photos/about", verbose_name = 'Второстепенное фотография')
    image1 = models.ImageField(blank=True, null=True, upload_to="data/photos/about", verbose_name = 'Фото 1')
    image2 = models.ImageField(blank=True, null=True, upload_to="data/photos/about", verbose_name = 'Фото 2')
    image3 = models.ImageField(blank=True, null=True, upload_to="data/photos/about", verbose_name = 'Фото 3')
    image4 = models.ImageField(blank=True, null=True, upload_to="data/photos/about", verbose_name = 'Фото 4')
    image5 = models.ImageField(blank=True, null=True, upload_to="data/photos/about", verbose_name = 'Фото 5')
    image6 = models.ImageField(blank=True, null=True, upload_to="data/photos/about", verbose_name = 'Фото 6')
    image7 = models.ImageField(blank=True, null=True, upload_to="data/photos/about", verbose_name = 'Фото 7')

    class Meta:
        managed = False
        db_table = 'abouts'
        verbose_name = 'О компании'
        verbose_name_plural = 'О компании'

    def __str__(self):
        return f'{self.id}'
