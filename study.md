+ Device
    1st Device Entrypoint
        Direct TCP/IP
            Host
            Port
            # The group-key here is Host+Port
        Controlled Serial over IP (Moxa) = TCP/IP
            Host
            Port
            Command Port
            # The group-key here is Host+Port
        Phone Line (Modem) => Modem over TCP/IP
            Phone Number
            # The group-key here is Phone Number
            # In this case the device needs to be assigned to a pool

                            []Pool
                                []Modem
                                    Id
                                    []Init AT
                                    ...

        # The group-keys the entrypoints
        # The group-keys solves also Data Concentrator and Gateway design!

        # Push mode?
        # - DLMS push TCP/IP is push

    Driver "Driver clbs DLMS Generic"         [List of installed drivers?]
           "Driver Landis+Gyr E650"
           "Service clbs DLMS Push"
           "Service clbs MQTT Push"
    Protocol "DLMS/LN over HDLC"              [List of protocols supported by the driver]                   ?? Is this true or we develop driver with single protocol yet configurable?
             "DLMS/SN over COSEM Wrapper"
             "IEC 62056-21 (IEC-61107)"
             "DLMS Push"
             "MQTT Push"
             # Not all protocols can be user via all CU types - do we want to check or allow all?
    Device
        # Static attributes
        Name "Meter 1"
             "12345 PRG"
        Serial Number
             "0012345678"
        # Dynamic attributes
        Protocol-related attributes         [List of attributes]
            HDLC Address
            IEC Address
            Push ID
            MQTT what-ever
