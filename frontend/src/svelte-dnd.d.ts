// takes care of types for dnd action until https://github.com/sveltejs/language-tools/issues/431 is resolved
declare type Item = import("svelte-dnd-action").Item;
declare type DndEvent<ItemType = Item> = import("svelte-dnd-action").DndEvent<ItemType>;
declare namespace svelteHTML {
    interface HTMLAttributes<T> {
        "on:consider"?: (event: CustomEvent<DndEvent<ItemType>> & {target: EventTarget & T}) => void;
        "on:finalize"?: (event: CustomEvent<DndEvent<ItemType>> & {target: EventTarget & T}) => void;
    }
}