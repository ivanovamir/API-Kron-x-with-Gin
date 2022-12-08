from django.contrib import admin

# Register your models here.

from .models import *


@admin.register(Abouts)
class AboutsAdmin(admin.ModelAdmin):
    list_display = ('id', )
    list_display_links = ('id', )

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(HeadOffices)
class HeadOfficesAdmin(admin.ModelAdmin):
    list_display = ('id', 'title', 'phone', 'mail')
    list_display_links = ('id', 'title', 'phone', 'mail')


@admin.register(Requisites)
class RequisitesAdmin(admin.ModelAdmin):
    list_display = ('id', 'company_name', 'address', 'inn', 'ogrn', 'rs', 'bank', 'ks', 'bik')
    list_display_links = ('id', 'company_name', 'address', 'inn', 'ogrn', 'rs', 'bank', 'ks', 'bik')
    
    def has_add_permission(self, request, obj=None):
        return False

    def has_delete_permission(self, request, obj=None):
        return False

