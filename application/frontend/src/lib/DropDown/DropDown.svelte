<!-- @migration-task Error while migrating Svelte code: `<thbody>` cannot be a child of `<table>`. `<table>` only allows these children: `<caption>`, `<colgroup>`, `<tbody>`, `<thead>`, `<tfoot>`, `<style>`, `<script>`, `<template>`. The browser will 'repair' the HTML (by moving, removing, or inserting elements) which breaks Svelte's assumptions about the structure of your components.
https://svelte.dev/e/node_invalid_placement -->
<script>
  import DropDownRow from "./dropDownRow.svelte";
</script>

<div class="dropDownContainer ">
  <table class="dropDownTable">
    <tbody>
      <DropDownRow
        rowTitle={"Break long lines"}
        rowImage={"log log-Break"}
        storeValue={"breakLines"}
        isFirst={true}
      />
      <DropDownRow
        rowTitle={"Log time in UTC"}
        rowImage={"log log-Clock"}
        storeValue={"UTCtime"}
      />
      <DropDownRow
        rowTitle={"Case insensitive search"}
        rowImage={"log log-Leters"}
        storeValue={"caseInSensitive"}
      />
      <DropDownRow
        rowTitle={"Transform JSON "}
        rowImage={"log log-Json"}
        storeValue={"transformJson"}
      />
    </tbody>
  </table>
</div>
