from django.urls import path

from cms.views import CmsGoogleMetricLink, CmsYandexMetricLink

urlpatterns = [
    path('google_metric/', CmsGoogleMetricLink, name = 'google_metric'),
    path('yandex_metric/', CmsYandexMetricLink, name = 'yandex_metric'),
]
