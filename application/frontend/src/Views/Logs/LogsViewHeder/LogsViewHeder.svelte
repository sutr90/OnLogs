<script>
  

  import Button from "../../../lib/Button/Button.svelte";
  import DropDown from "../../../lib/DropDown/DropDown.svelte";
  import { clickOutside } from "../../../lib/OutsideClicker/OutsideClicker.js";
  /**
   * @typedef {Object} Props
   * @property {string} [searchText]
   */

  /** @type {Props} */
  let { searchText = $bindable("") } = $props();
  let dropDownIsVisible = $state(false);
  let isSearchVIsible = $state(false);
  let timer;
  const debounce = (v) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      searchText = v;
    }, 750);
  };

  function dropDownToggle() {
    dropDownIsVisible = !dropDownIsVisible;
  }
  function handleClickOutside() {
    if (dropDownIsVisible) {
      dropDownIsVisible = false;
    }
  }
</script>

<div id="top-line">
  <h2 class="header">Service logs</h2>
  <p class="header">recent at bottom</p>
  <!-- <button class="header show">
        <div class="icon">
            <i class={"log log-Eye"} />
        </div>
    </button> -->
  {#if !isSearchVIsible}
    <div
      style:position={"relative"}
      use:clickOutside
      onclick_outside={handleClickOutside}
    >
      <Button
        icon={"log log-Eye"}
        minWidth={55}
        minHeight={40}
        iconHeight={20}
        CB={dropDownToggle}
      />
      {#if dropDownIsVisible}
        <DropDown />
      {/if}
    </div>
    <div class="filterButtonContainer visuallyHidden">
      <Button
        icon={"log log-Filter"}
        minWidth={55}
        minHeight={40}
        iconHeight={20}
        CB={() => {}}
      />
    </div>
  {/if}
  <div class="searchButtonContainer">
    <Button
      icon={"log log-Search"}
      minWidth={55}
      minHeight={40}
      iconHeight={20}
      CB={() => {
        isSearchVIsible = !isSearchVIsible;
      }}
    />
  </div>
  <div class="header search {!isSearchVIsible && 'hidden'}">
    {#if !searchText}<div class="searchIcoContainer">
        <i class={"log log-Search"}></i>
      </div>{/if}
    <input
      type="text"
      oninput={(e) => {
        searchText = e.target.value;
      }}
      placeholder="Search"
    />
  </div>
</div>
