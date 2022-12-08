from django.contrib import admin

# Register your models here.

from .models import *


@admin.register(FeedbackCalls)
class FeedbackCallsAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'phone')
    list_display_links = ('id', 'name', 'phone')

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(FeedbackSelections)
class FeedbackSelectionsAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'phone', 'email')
    list_display_links = ('id', 'name', 'phone', 'email')
    
    def has_add_permission(self, request, obj=None):
        return False


@admin.register(Supports)
class SupportsAdmin(admin.ModelAdmin):
    list_display = ('id', 'name')
    list_display_links = ('id', 'name')
    
    def has_add_permission(self, request, obj=None):
        return False