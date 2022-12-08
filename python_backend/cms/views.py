from django.shortcuts import redirect

from .models import Metrics


def CmsGoogleMetricLink(request):
    link = Metrics.objects.get(id=1)
    field_name_val = getattr(link, "google_metric")
    return redirect(field_name_val)


def CmsYandexMetricLink(request):
    link = Metrics.objects.get(id=1)
    field_name_val = getattr(link, "yandex_metric")
    return redirect(field_name_val)
