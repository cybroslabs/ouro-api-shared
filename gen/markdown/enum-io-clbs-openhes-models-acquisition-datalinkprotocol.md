# Enum: io.clbs.openhes.models.acquisition.DataLinkProtocol

Data link protocols

## Options

| Value | Description |
| --- | --- |
| LINKPROTO_IEC_62056_21 | The VDEW (IEC 62056-21, IEC-61107) protocol. In combination with DLMS protocol the driver initiates the communication by IEC and switches to the mode E to the HDLC+DLMS protocol. Supports addressing = multiple devices on the same line. |
| LINKPROTO_HDLC | The HDLC (ISO/IEC-3309) framing. It can be used for various application protocols, such as DLMS or MODBUS. Supports client/server addressing = multiple devices on the same line. |
| LINKPROTO_COSEM_WRAPPER | The COSEM wrapper. It can be used for DLMS application protocol. Supports client/server addressing = multiple devices on the same line. |
| LINKPROTO_MODBUS | The Modbus protocol. It shall be used for Modbus application protocol where no other data link layer, such as HDLC, is used. |
| LINKPROTO_MBUS | The M-Bus protocol. It shall be used for M-Bus application protocol. |
| LINKPROTO_VIKTOR | The Viktor protocol. It is a proprietary protocol used by Viktor-based devices, such as DC450 Vitkor. |
| LINKPROTO_NOT_APPLICABLE | The data link protocol is not applicable. It's useful for listening communication type. |
