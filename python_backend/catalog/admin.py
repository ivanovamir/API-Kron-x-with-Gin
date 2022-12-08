from django.contrib import admin

from .models import *
from .views import validate_file_extension


class CartProductsInline(admin.TabularInline):
    model = CartProducts
    extra = 1
    verbose_name = "Продукт корзины"
    verbose_name_plural = "Продукты корзины"

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False

    def has_change_permission(self, request, obj=None):
        return False


class OptionOfDeliveryAndPaymentsInline(admin.TabularInline):
    model = OptionOfDeliveryAndPayments
    extra = 1
    verbose_name = "Доставка"
    verbose_name_plural = "Доставка"

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False

    def has_change_permission(self, request, obj=None):
        return False


@admin.register(Categories)
class CategoriesAdmin(admin.ModelAdmin):
    list_display = ('id', 'title')
    list_display_links = ('id', 'title')
    list_per_page = 15
    search_fields = ('id', 'title')
    search_help_text = 'Введите id или название'
    list_filter = ('title',)


@admin.register(Products)
class ProductsAdmin(admin.ModelAdmin):
    list_display = ('id', 'title', 'category', 'price')
    list_display_links = ('id', 'title', 'price', 'category')
    list_per_page = 15
    search_fields = ('id', 'title', 'price')
    search_help_text = 'Введите id,название или цену'
    list_filter = ('category', 'can_to_view',)


@admin.register(Orders)
class OrdersAdmin(admin.ModelAdmin):
    list_display = ('id', 'created_at', 'order_status')
    list_display_links = ('id', 'created_at', 'order_status')
    list_per_page = 15
    readonly_fields = ('created_at', 'note', 'form')
    search_fields = ('id', 'created_at')
    search_help_text = 'Введите id заказа или дату'
    list_filter = ('order_status',)
    inlines = (CartProductsInline, OptionOfDeliveryAndPaymentsInline)

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(Users)
class UsersAdmin(admin.ModelAdmin):
    list_display = ('id', 'Категория', 'email', 'inn', 'phone',)
    list_display_links = ('id', 'email', 'inn')
    list_per_page = 15
    search_fields = ('phone', 'inn')
    search_help_text = 'Телефон или ИНН'
    readonly_fields = ('created_at',)

    def Категория(self, obj):
        if obj.inn == "":
            return f'Физ. лицо'
        else:
            return 'Юр. лицо'

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(OrderStatuses)
class OrderStatusesAdmin(admin.ModelAdmin):
    list_display = ('id', 'name',)
    list_display_links = ('id', 'name',)
    list_per_page = 15


@admin.register(Catalogs)
class CatalogsAdmin(admin.ModelAdmin):
    list_display = ('id', 'title',)
    list_display_links = ('id', 'title',)
    list_per_page = 15
    search_fields = ('title',)
    search_help_text = 'Название каталога'
