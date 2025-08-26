# Model: io.clbs.openhes.models.common.ListSelector

Defines the selection criteria for listing objects, including pagination, sorting, filtering, and field selection.

## Fields

| Field | Information |
| --- | --- |
| pageSize | <b>Type:</b> `uint32`<br><b>Description:</b><br>The number of items per page. |
| offset | <b>Type:</b> `uint32`<br><b>Description:</b><br>The zero-based offset of the first item in the response.<br><b>Values:</b> Any non-negative integer<br><b>Example:</b> 15 |
| sortBy | <b>Type:</b> [`io.clbs.openhes.models.common.ListSelectorSortBy`](model-io-clbs-openhes-models-common-listselectorsortby.md)<br><b>Description:</b><br>The sorting criteria. |
| filterBy | <b>Type:</b> [`io.clbs.openhes.models.common.ListSelectorFilterBy`](model-io-clbs-openhes-models-common-listselectorfilterby.md)<br><b>Description:</b><br>The filtering criteria. |
| fields | <b>Type:</b> `string`<br><b>Description:</b><br>FIXME: This needs to be designed properly.<br><br>Optional list of additional fields to include in the response. |

