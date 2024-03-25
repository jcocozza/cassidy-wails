<script lang="ts">
    import { onMount } from 'svelte';
    import EquipmentViewer from '$lib/components/equipment/EquipmentViewer.svelte';
    import type { model } from '../../../../wailsjs/go/models';
    import { List, ListEquipmentTypes } from '../../../../wailsjs/go/controllers/EquipmentHandler';

    export let usr: model.User;
    let equipment_list: model.Equipment[] = [];
    let equipment_type_list: model.EquipmentType[] = [];
    onMount(async () => {
        equipment_list = await List(usr)
        equipment_type_list = await ListEquipmentTypes()
    })
</script>
<EquipmentViewer bind:usr={usr} bind:equipment_list={equipment_list} equipment_type_list={equipment_type_list}></EquipmentViewer>
