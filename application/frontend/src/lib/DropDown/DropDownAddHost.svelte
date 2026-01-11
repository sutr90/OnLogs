<!-- @migration-task Error while migrating Svelte code: `<thbody>` cannot be a child of `<table>`. `<table>` only allows these children: `<caption>`, `<colgroup>`, `<tbody>`, `<thead>`, `<tfoot>`, `<style>`, `<script>`, `<template>`. The browser will 'repair' the HTML (by moving, removing, or inserting elements) which breaks Svelte's assumptions about the structure of your components.
https://svelte.dev/e/node_invalid_placement -->
<script>
  import DropDownRow from "./dropDownRow.svelte";
  import {
    snipetModalIsVisible,
    addHostMenuIsVisible,
  } from "../../Stores/stores.js";
</script>

<div class="dropDownContainer addHost " id="addHost">
  <table class="dropDownTable">
    <tbody>
      <DropDownRow
        rowTitle={"Create services group"}
        rowImage={"log log-Group"}
        storeValue={""}
        isFirst={true}
        disableCheckbox={true}
        titleCallBack={() => {}}
      />
      <DropDownRow
        rowTitle={"Connect new host"}
        rowImage={"log log-Server"}
        storeValue={""}
        disableCheckbox={true}
        titleCallBack={() => {
          snipetModalIsVisible.set(true);
          addHostMenuIsVisible.set(false);
        }}
      />
    </tbody>
  </table>
</div>
