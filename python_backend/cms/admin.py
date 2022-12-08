from django.contrib import admin

from .models import *

admin.site.site_url = None


@admin.register(Features)
class FeaturesAdmin(admin.ModelAdmin):
    list_display = ('id', 'header')
    list_display_links = ('id', 'header')
    list_per_page = 15


@admin.register(Sliders)
class SlidersAdmin(admin.ModelAdmin):
    list_display = ('id', 'main_text', 'upper_text', 'down_text')
    list_display_links = ('id', 'main_text', 'upper_text', 'down_text')
    list_per_page = 15
    search_fields = ('id',)
    search_help_text = 'Введите id'


@admin.register(Services)
class ServicesAdmin(admin.ModelAdmin):
    list_display = ('id', 'title',)
    list_display_links = ('id', 'title',)


@admin.register(Emails)
class EmailsAdmin(admin.ModelAdmin):
    list_display = ('id', 'address',)
    list_display_links = ('id', 'address',)

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(Metrics)
class MetricsAdmin(admin.ModelAdmin):
    list_display = ('google_metric', 'yandex_metric',)
    list_display_links = ('google_metric', 'yandex_metric',)

    def has_delete_permission(self, request, obj=None):
        return False

    def has_add_permission(self, request, obj=None):
        return False
