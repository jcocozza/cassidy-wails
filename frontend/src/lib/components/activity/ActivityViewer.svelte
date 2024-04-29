<script lang="ts">
    import type { model } from "$lib/wailsjs/go/models";
    import { LineLayer, MapLibre, GeoJSON } from 'svelte-maplibre';
    import type { Feature } from 'geojson';
    import { decode } from "@mapbox/polyline";
    import { onMount } from "svelte";
    import { ConvertDuration } from "$lib/model/date";
    export let activity: model.Activity

    let data: Feature;
    let coords: [number, number][];
    let map_center: [number, number];

    // convert google polyline string to coordinate pairs
    //
    // I'm not implementing this in go (yet), because it doesn't seem to be worth the hassle.
    function PolylineToCoords(polyline: string): [number, number][] {
        const decodedCoordinates: [number, number][] = decode(polyline)
        // in the dumbest thing ever, map libre reads coords long, lat
        // while the decode is lat, long
        decodedCoordinates.forEach((coords) => {
            coords.reverse()
        })

        return decodedCoordinates
    }

    onMount(() => {
        coords = PolylineToCoords(activity.map)
        map_center = coords[0]
        data = {
            type: 'Feature',
            properties: {
                name: 'mapppp'
            },
            geometry: {
                type: 'Polygon',
                coordinates: [coords]
            }
        }
    })

</script>

<div class="container">
    <h1>{activity.activity_type?.name} - {activity.name}</h1>
    <h3>{activity.completed?.distance?.length} {activity.completed?.distance?.unit}  {activity.completed?.pace}</h3>
    {#each activity.type_subtype_list as subtype}
        <span style="margin-right: 1px; margin-bottom: 1px;" class="badge rounded-pill text-bg-secondary tag">{subtype.activity_subtype.name}</span>
    {/each}

    <div class="row">
        <p>{activity.description}</p>
    </div>

    <div class="row">

        <div class="col-lg-4">
            <div class="row">
                <p>{activity.notes}</p>
            </div>

            <div class="row">
                <table class="tbl">
                    <tr>
                        <th>Planned</th>
                        <th>Completed</th>
                    </tr>

                    <tr>
                        <td>{activity.planned?.distance?.length} {activity.planned?.distance?.unit}</td>
                        <td>{activity.completed?.distance?.length} {activity.completed?.distance?.unit}</td>
                    </tr>
                    <tr>
                        <td>{ConvertDuration(activity.planned?.duration)}</td>
                        <td>{ConvertDuration(activity.completed?.duration)}</td>
                    </tr>
                    <tr>
                        <td>{activity.planned?.vertical?.length} {activity.planned?.vertical?.unit}</td>
                        <td>{activity.completed?.vertical?.length} {activity.completed?.vertical?.unit}</td>
                    </tr>
                </table>
            </div>

            <div class="row">
                {#each activity.equipment_list as equipment}
                    {equipment.equipment?.name}: {equipment.assigned_mileage?.length} {equipment.assigned_mileage?.unit}
                {/each}
            </div>
        </div>
        <!-- Map goes here -->
        <div class="col-lg-8" style="min-height: 70vh;">
            <MapLibre
                    center={map_center}
                    style="https://basemaps.cartocdn.com/gl/positron-gl-style/style.json"
                    class="map"
                    standardControls
                    zoom={15}
                    hash
                    >
                    <GeoJSON id="mapppp" {data}>
                        <LineLayer
                        layout={{ 'line-cap': 'round', 'line-join': 'round' }}
                        paint={{ 'line-color': '#0000FF', 'line-width': 3 }}
                        />
                    </GeoJSON>
                </MapLibre>
        </div>
    </div>
</div>

<style>
  :global(.map) {
    height: 100%;
    width: 100%;
  }

  .tbl {
    border: 1px solid lightgray;
  }

</style>
