<script lang="ts">
    import type { model } from "../../../../wailsjs/go/models";

    export let activity: model.Activity;
</script>

<div class="container">

    <div class="row">
        <div class="col">
            {activity.notes}
        </div>
    </div>

</div>