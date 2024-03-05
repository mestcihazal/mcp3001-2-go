# `mcp3001-2` modular resource

The module provides analog-to-digital conversion capabilities for MCP3001 and MCP3002 SPI ADCs. Written in Go, board agnostic. Tested on Raspberry Pi and Jetson Orin.

The values you get from the readings will be proportional to the voltage and you will need to interpret it depending on which sensor you hook up to the MCP300x.

You can add as many analog sensors as your MCP300x allows and get readings from them concurrently (this depends on how many channels it has, so for MCP3001 that is one channel, and for MCP3002 that is two).

## Build and run

To use this module, follow the instructions to [add a module from the Viam Registry](https://docs.viam.com/registry/configure/#add-a-modular-resource-from-the-viam-registry) and select the `hazalmestci:sensor:mcp3001-2-go` model from the [`mcp3001-2` module](https://app.viam.com/module/hazalmestci/mcp3001-2).

## Configure your `mcp3001-2`

> [!NOTE]
> Before configuring your `mcp3001-2`, you must [create a machine](https://docs.viam.com/manage/fleet/machines/#add-a-new-machine).

Navigate to the **Config** tab of your machine's page in [the Viam app](https://app.viam.com/).
Click on the **Components** subtab and click **Create component**.
Select the `sensor` type, then select the `hazalmestci:sensor:mcp3001-2-go` model.
Click **Add module**, then enter a name for your sensor and click **Create**.

On the new component panel, copy and paste the following attribute template into your sensor’s **Attributes** box:

> [!NOTE]
> For more information, see [Configure a Machine](https://docs.viam.com/manage/configuration/).

```json
{
    "chip_select": "0",
    "spi_bus": "0",
    "pins": {
        "moisture": 0,
        "temperature": 1
    }
}
```

Save your config.

### Attributes

The following attributes are available for a `mcp3001-2` sensor:

| Name    | Type   | Inclusion    | Description |
| ------- | ------ | ------------ | ----------- |
| `chip_select` | string | **Required** | The `chip_select`` pin you are using. |
| `spi_bus` | string | **Required** | the `spi_bus` you are using. |
| `pins` | string | **Required** | The pins you are using for moisture, temperature, and humidity. |

### Example configuration

```json
{
    "name": "my-mcp3001",
    "model": "hazalmestci:sensor:mcp3001-2-go",
    "type": "sensor",
    "namespace": "rdk",
    "attributes": {
        "chip_select": "0",
        "spi_bus": "0",
        "pins": {
            "moisture": 0,
            "temperature": 1
        }
    },
    "depends_on": []
}
```

## Local Development

To use the `mcp3001-2` module with a local install, clone this repository to your machine’s computer, navigate to the `module` directory, and run:

```go
go build
```

On your robot’s page in the [Viam app](https://app.viam.com/), enter
the [module’s executable path](/registry/create/#prepare-the-module-for-execution), then click **Add module**.
The name must use only lowercase characters.
Then, click **Save config**.

## Next Steps

1. To test your sensor, go to the [**Control** tab](https://docs.viam.com/manage/fleet/robots/#control) and test that you are getting readings.
2. Once you can obtain your readings, configure the data manager to [capture](https://docs.viam.com/data/capture/) and [sync](https://docs.viam.com/data/cloud-sync/) the data from all of your machines.
3. To retrieve data captured with the data manager, you can [query data with SQL or MQL](https://docs.viam.com/data/query/) or [visualize it with tools like Grafana](https://docs.viam.com/data/visualize/).

## License

Copyright 2021-2023 Viam Inc. <br>
Apache 2.0
