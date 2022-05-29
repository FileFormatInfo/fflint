---
title: Flags
---


## Common Flags

These flags can be used with any [command](/commands/index.html).

<table class="table table-striped table-bordered">
    <tr>
        <th>Flag</th>
        <th>Type</th>
        <th>Default</th>
        <th>Description</th>
    </tr>
{%- for option in site.data.common_flags %}
    <tr>
        <td>{{option.name}}</td>
        <td>{{option.type}}</td>
        <td>{{option.default_value}}</td>
        <td>{{option.usage | replace: "|", "&#x7c;" | markdownify | remove: '<p>' | remove: '</p>' }}</td>
    </tr>
{% endfor %}
</table>

## Custom Flag Types

### Range

Ranges specify acceptable minimum and maximum values.

* The minimum and maximum values are separated by a dash.
* If either the minimum or maximum values are missing, there is no minimum/maximum.
* If there is just one value (and no dash), that is the only acceptable value (i.e. it is both the minimum and maximum)
* Ranges are inclusive.

<!-- LATER: examples -->

### Ratio

Ratios specify the acceptable proportions between two values.

* N:M
* Decimal number

<!-- LATER: examples -->
