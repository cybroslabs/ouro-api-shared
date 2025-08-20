# Enum: io.clbs.openhes.models.acquisition.DataLinkProtocol

Data link protocols

## Options

| Value | Description |
| --- | --- |
| LINKPROTO_IEC_62056_21 | The VDEW (IEC 62056-21, IEC-61107) protocol. When combined with the DLMS protocol, the driver initiates communication via IEC and switches to the mode E (HDLC+DLMS) protocol. Supports multiple devices on the same line. |
| LINKPROTO_HDLC | The HDLC (ISO/IEC-3309) framing. Used with various application protocols, such as DLMS or MODBUS. Supports client/server addressing for multiple devices on the same line. |
| LINKPROTO_COSEM_WRAPPER | The COSEM wrapper. Used with the DLMS application protocol. Supports client/server addressing for multiple devices on the same line. |
| LINKPROTO_MODBUS | The Modbus protocol. Used with the Modbus application protocol where no other data link , such as HDLC is used. |
| LINKPROTO_MBUS | The M-Bus protocol. Used with the M-Bus application protocol. |
| LINKPROTO_VIKTOR | The Viktor protocol. A proprietary protocol used by Viktor-based devices, such as DC450 Vitkor. |
| LINKPROTO_NOT_APPLICABLE | No data link protocol is applicable. Used with listening communication types. |
