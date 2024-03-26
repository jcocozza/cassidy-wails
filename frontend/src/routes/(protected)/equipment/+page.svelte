<script lang="ts">
    import { onMount } from 'svelte';
    import EquipmentViewer from '$lib/components/equipment/EquipmentViewer.svelte';
    import type { model } from '$lib/wailsjs/go/models';
    import { List, ListEquipmentTypes } from '$lib/wailsjs/go/controllers/EquipmentHandler';

    let equipment_list: model.Equipment[] = [];
    let equipment_type_list: model.EquipmentType[] = [];

    onMount(async () => {
        equipment_list = await List()
        equipment_type_list = await ListEquipmentTypes()
    })
</script>

<EquipmentViewer bind:equipment_list={equipment_list} equipment_type_list={equipment_type_list}></EquipmentViewer>