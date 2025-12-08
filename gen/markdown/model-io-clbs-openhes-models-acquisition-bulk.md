# Model: io.clbs.openhes.models.acquisition.Bulk

Defines a bulk operation model for executing actions across multiple devices concurrently.
 Bulks enable efficient mass operations such as data collection, firmware updates, or configuration changes
 across hundreds or thousands of devices. Each bulk creates individual jobs per device.

## Fields

| Field | Information |
| --- | --- |
| spec | <b>Type:</b> [`io.clbs.openhes.models.acquisition.BulkSpec`](model-io-clbs-openhes-models-acquisition-bulkspec.md)<br><b>Description:</b><br>The bulk specification defining target devices and actions. |
| status | <b>Type:</b> [`io.clbs.openhes.models.acquisition.BulkStatus`](model-io-clbs-openhes-models-acquisition-bulkstatus.md)<br><b>Description:</b><br>The current status tracking job counts and completion state. |
| metadata | <b>Type:</b> [`io.clbs.openhes.models.common.MetadataFields`](model-io-clbs-openhes-models-common-metadatafields.md)<br><b>Description:</b><br>Metadata including id, name, generation, user-managed fields, and system-managed fields. |

