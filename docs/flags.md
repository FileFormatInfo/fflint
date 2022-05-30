---
title: Flags
---


## Common Flags

These flags can be used with any [command](/commands/index.html).

<div class="table-responsive">
    <table class="table table-striped table-bordered">
        <tr>
            <th>Flag</th>
            <th class="d-none d-md-table-cell">Type</th>
            <th>Default</th>
            <th>Description</th>
        </tr>
{%- for option in site.data.common_flags %}
        <tr>
            <td>{{option.name}}</td>
            <td class="d-none d-md-table-cell">{{option.type}}</td>
            <td>{{option.default_value}}</td>
            <td>{{option.usage | replace: "|", "&#x7c;" | markdownify | remove: '<p>' | remove: '</p>' }}</td>
        </tr>
{% endfor %}
    </table>
</div>

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
