from django.contrib import messages
from django.db import models
from django.http import HttpResponseRedirect
from .views import validate_file_extension


class Categories(models.Model):
    id = models.BigAutoField(primary_key=True, verbose_name='Id')
    title = models.CharField(max_length=100, blank=True, null=True, verbose_name='Имя категории')
    image = models.ImageField(blank=True, null=True, upload_to="data/photos/categories", verbose_name='Фото категории')

    class Meta:
        managed = False
        db_table = 'categories'
        verbose_name = 'Категория'
        verbose_name_plural = 'Категории'
        ordering = ['id']

    def __str__(self):
        return f'Категория - {self.title}'


class Products(models.Model):
    category = models.ForeignKey(Categories, models.DO_NOTHING, db_column='category', blank=True, null=True,
                                 verbose_name='Категория')
    title = models.CharField(max_length=250, blank=True, null=True, verbose_name='Название')
    vendor_code = models.CharField(max_length=100, blank=True, null=True, verbose_name='Артикул')
    code1c = models.CharField(max_length=100, blank=True, null=True, verbose_name='Код 1с')
    description = models.TextField(blank=True, null=True, verbose_name='Описание')
    price = models.DecimalField(max_digits=15, decimal_places=2, blank=True, null=True, verbose_name='Цена')
    image_original = models.ImageField(null=True, blank=True, upload_to="data/photos/products/products_original",
                                       verbose_name='Главная фотография')
    image128 = models.ImageField(null=True, blank=True, upload_to="data/photos/products/products_128",
                                  verbose_name='Фотография 128px')
    image432 = models.ImageField(null=True, blank=True, upload_to="data/photos/products/products_432",
                                  verbose_name='Фотография 432px')
    can_to_view = models.BooleanField(blank=False, null=False, verbose_name='Разрешение на публикацию',
                                      help_text='На сайте отоброжаются только те продукты, где данное поле принимает значение "Да" | По умолчанию при добавлении продукта данное поле имеет положительно значение')

    class Meta:
        managed = False
        db_table = 'products'
        verbose_name = 'Товар'
        verbose_name_plural = 'Товары'
        ordering = ['id']

    def __str__(self):
        return f'Название- {self.title} | {self.category}'


class Users(models.Model):
    id = models.BigAutoField(primary_key=True)
    created_at = models.DateTimeField(blank=True, null=True, verbose_name='Дата создания')
    inn = models.CharField(blank=True, null=True, max_length=50, verbose_name='ИНН')
    email = models.CharField(blank=True, null=True, max_length=50, verbose_name='Email')
    phone = models.CharField(blank=True, null=True, max_length=50, verbose_name='Телефон')

    class Meta:
        managed = False
        db_table = 'users'
        verbose_name = 'Форма покупателя'
        verbose_name_plural = 'Формы покупателей'
        ordering = ['id']

    def __str__(self):
        if self.inn == "":
            return f'Физ. лицо: Номер телефона - {self.phone}'
        else:
            return f'Юр. лицо: Email - {self.email} | ИНН - {self.inn}'


class Carts(models.Model):
    id = models.BigAutoField(primary_key=True)
    user = models.ForeignKey(Users, models.DO_NOTHING, db_column="user_id", blank=True, null=True,
                             verbose_name='Форма покупателя')
    in_order = models.BooleanField(blank=True, null=True, verbose_name='В заказе', editable=False)
    order = models.ForeignKey('Orders', models.DO_NOTHING, db_column="order_id", blank=True, null=True,
                              verbose_name='Форма покупателя')

    class Meta:
        managed = False
        db_table = 'carts'
        verbose_name = 'Корзина'
        verbose_name_plural = 'Корзины'

    def __str__(self):
        return f'Номер корзины - {self.id}'


class CartProducts(models.Model):
    id = models.BigAutoField(primary_key=True)
    product = models.ForeignKey('Products', models.DO_NOTHING, blank=True, null=True, verbose_name='Продукт')
    count = models.BigIntegerField(blank=True, null=True, verbose_name='Кол-во')
    order = models.ForeignKey('Orders', models.DO_NOTHING, db_column='order_id', blank=True, null=True,
                              verbose_name='Заказ')
    total_price = models.DecimalField(max_digits=65535, decimal_places=65535, blank=True, null=True,
                                      verbose_name='Общая сумма')

    class Meta:
        managed = False
        db_table = 'cart_products'
        verbose_name = 'Продукт корзины'
        verbose_name_plural = 'Продукты корзин'
        ordering = ['-id']

    def __str__(self):
        return f'Номер корзины - {self.id} | Кол-во - {self.count}'


class Favourites(models.Model):
    id = models.BigAutoField(primary_key=True)
    user = models.ForeignKey(Users, models.DO_NOTHING, db_column="user_id", blank=True, null=True,
                             verbose_name="Форма покупателя")

    class Meta:
        managed = False
        db_table = 'favourites'
        verbose_name = 'Корзина избранных'
        verbose_name_plural = 'Корзины избранных'

    def __str__(self):
        return f'Корзина избранных - {self.id} | Пользователь - {self.user}'


class FavouritesProducts(models.Model):
    id = models.BigAutoField(primary_key=True)
    product = models.ForeignKey('Products', models.DO_NOTHING, blank=True, null=True, verbose_name='Продукт')

    class Meta:
        managed = False
        db_table = 'favourites_products'
        verbose_name = 'Избранный продукт'
        verbose_name_plural = 'Избранные продукты'

    def __str__(self):
        return f'Продукт - {self.product}'


class Orders(models.Model):
    id = models.BigAutoField(primary_key=True, verbose_name='Номер заказа')
    created_at = models.DateTimeField(blank=True, null=True, verbose_name='Дата создания')
    note = models.TextField(blank=True, null=True, verbose_name='Примечание к заказу')
    order_status = models.ForeignKey('OrderStatuses', models.DO_NOTHING, blank=True, null=True,
                                     verbose_name='Статус заказа')
    user = models.ForeignKey(Users, models.DO_NOTHING, db_column="user_id", blank=True, null=True,
                             verbose_name='Покупатель')

    class Meta:
        managed = False
        db_table = 'orders'
        verbose_name = 'Заказ'
        verbose_name_plural = 'Заказы'
        ordering = ['-created_at']

    def __str__(self):
        return f'Номер заказа - {self.id} | Дата создания - {self.created_at}'


class OrderStatuses(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.CharField(max_length=50, blank=True, null=True, verbose_name='Имя')

    class Meta:
        managed = False
        db_table = 'order_statuses'
        verbose_name = 'Статус заказа'
        verbose_name_plural = 'Статусы заказов'

    def __str__(self):
        return f'{self.name}'


class OptionOfDeliveryAndPayments(models.Model):
    id = models.BigAutoField(primary_key=True)
    city = models.TextField(blank=True, null=True, verbose_name='Город')
    street = models.TextField(blank=True, null=True, verbose_name='Улица')
    house = models.TextField(blank=True, null=True, verbose_name='Дом')
    frame = models.TextField(blank=True, null=True, verbose_name='Корпус')
    entrance = models.TextField(blank=True, null=True, verbose_name='Подъезд')
    apartment_office = models.TextField(blank=True, null=True, verbose_name='Квартира/Оффис')
    order = models.ForeignKey('Orders', models.DO_NOTHING, blank=True, null=True, verbose_name='Заказ')

    class Meta:
        managed = False
        db_table = 'option_of_delivery_and_payments'
        verbose_name = 'Информация о доставке'
        verbose_name_plural = 'Информация о доставке'


class Catalogs(models.Model):
    id = models.BigAutoField(primary_key=True)
    title = models.TextField(blank=True, null=True, verbose_name='Имя каталога')
    file = models.FileField(upload_to='data/catalogs', validators=[validate_file_extension], verbose_name='Файл')

    class Meta:
        managed = False
        db_table = 'catalogs'
        verbose_name = 'Каталог'
        verbose_name_plural = 'Каталоги'

    def upload_pdf_file(self, request):
        if self.title[-3:-1] != "PDF":
            messages.warning(request, 'Файл должен быть формата PDF')
            return HttpResponseRedirect("/")

    def __str__(self):
        return f'Каталог - {self.title}'
