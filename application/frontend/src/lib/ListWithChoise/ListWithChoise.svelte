<script>
  import { run } from 'svelte/legacy';

  // @ts-nocheck
  import { onMount } from "svelte";

  let sortedData = $state([]);
  import { navigate } from "svelte-routing";
  let openStopedServIndexes = $state([]);
  let chosenElSettings = $state("");
  import {
    activeMenuOption,
    lastChosenHost,
    lastChosenService,
    listScrollIsVisible,
  } from "../../Stores/stores.js";
  import FetchApi from "../../utils/fetch.js";
  import { changeKey } from "../../utils/changeKey.js";
  /**
   * @typedef {Object} Props
   * @property {any} [listData]
   * @property {any} [openHeaderIndexs]
   * @property {string} [activeElementName]
   * @property {string} [customListClass]
   * @property {string} [customListElClass]
   * @property {string} [customActiveElClass]
   * @property {string} [headerButton]
   * @property {string} [listElementButton]
   */

  /** @type {Props} */
  let {
    listData = $bindable([]),
    openHeaderIndexs = $bindable([0]),
    activeElementName = $bindable(""),
    customListClass = "",
    customListElClass = "",
    customActiveElClass = "",
    headerButton = "",
    listElementButton = ""
  } = $props();
  let initialVisitcounter = $state(0);

  const fetchApi = new FetchApi();

  run(() => {
    {
      if (!initialVisitcounter && !activeElementName) {
        const chosenHost =
          $lastChosenHost || (sortedData[0] && sortedData[0].host);
        const chosenService =
          $lastChosenService ||
          (sortedData[0] && sortedData[0].services[0].serviceName);
        activeElementName = sortedData[0] && `${chosenHost}-${chosenService}`;

        lastChosenHost.set(chosenHost);
        lastChosenService.set(chosenService);

        navigate(
          `${changeKey}/${
            location.pathname.includes("stats") ? "stats" : "view"
          }/${chosenHost}/${chosenService}`,
          {
            replace: true,
          }
        );
      }
    }
  });

  run(() => {
    sortedData = listData.map((h) => {
      let activeServices = h.services
        .filter((s) => {
          return !s.isDisabled;
        })
        .sort(function (a, b) {
          if (a.isFavorite < b.isFavorite) {
            return 1;
          }
          if (a.isFavorite > b.isFavorite) {
            return -1;
          }

          return 0;
        });
      let inActiveServices = h.services
        .filter((s) => {
          return s.isDisabled;
        })
        .sort(function (a, b) {
          if (a.isFavorite < b.isFavorite) {
            return 1;
          }
          if (a.isFavorite < b.isFavorite) {
            return -1;
          }

          return 0;
        });
      let newHost = {
        ...h,
        services: [...activeServices, ...inActiveServices],
      };
      return newHost;
    });
  });

  run(() => {
    if (!initialVisitcounter) {
      const openedStopedServiceIndex = sortedData
        .map((e) => {
          return e.services.map((s) => s.isDisabled && s.serviceName);
        })
        .findIndex((el) => {
          return el.includes($lastChosenService);
        });

      if (openedStopedServiceIndex !== -1) {
        openStopedServIndexes.push(openedStopedServiceIndex);
        openStopedServIndexes = [...new Set(openStopedServIndexes)];
      }
    }
  });

  async function favoriteToggle(host, service) {
    $lastChosenHost, $lastChosenService;

    const hostIndex = listData.findIndex((h) => h.host === host);
    const serviceIndex = listData[hostIndex].services.findIndex((s) => {
      return s.serviceName === service;
    });
    if (hostIndex !== -1) {
      if (listData[hostIndex].services[serviceIndex]) {
        listData[hostIndex].services[serviceIndex].isFavorite =
          !listData[hostIndex].services[serviceIndex].isFavorite;
      }
    }

    const data = await fetchApi.changeFavorite(host, service);
  }

  function toggleSublistVisible(i) {
    if (openHeaderIndexs.includes(i)) {
      openHeaderIndexs = openHeaderIndexs.filter((e) => e !== i);
    } else {
      openHeaderIndexs = [...openHeaderIndexs, i];
    }
  }
  function toggleArchivedVisible(i) {
    if (openStopedServIndexes.includes(i)) {
      openStopedServIndexes = openStopedServIndexes.filter((e) => e !== i);
    } else {
      openStopedServIndexes = [...openStopedServIndexes, i];
    }
  }

  function choseSublistEl(firstEl, secondEl) {
    activeElementName = `${firstEl.trim()}-${secondEl.trim()}`;

    navigate(
      `${changeKey}/${
        location.pathname.includes("stats") ? "stats" : "view"
      }/${firstEl}/${secondEl}`,
      {
        replace: true,
      }
    );
  }
  function choseInitialHost() {
    listData.forEach((h, i) => {
      if (h.host === $lastChosenHost) {
        openHeaderIndexs = [i, ...openHeaderIndexs];
      }
    });
  }

  run(() => {
    listData && choseInitialHost();
  });
</script>

<div class="listWithChoise {$listScrollIsVisible ? 'active' : ''}">
  <ul class={customListClass}>
    {#each sortedData as listEl, index}
      <li class="listElement">
        <div
          class="hostHeader clickable"
          onclick={({ target }) => {
            if (target.id !== headerButton) {
              toggleSublistVisible(index);
            }
          }}
        >
          <div>
            <i class="log log-Server"></i>
          </div>
          <p class="hostName">
            {listEl.host}
          </p>
          {#if headerButton}<div
              class="headerButton "
              id={headerButton}
              onclick={() => {
                console.log("clicable");
              }}
            >
              <i class="log log-{headerButton}" id={headerButton}></i>
            </div>{/if}
        </div>
        <div
          class="dropDownList {openHeaderIndexs.includes(index)
            ? ''
            : 'visuallyHidden'}"
        >
          <ul class="activeServices">
            {#each listEl.services as service, i}
              {#if !service.isDisabled}<li
                  class="serviceListItem  "
                  id={listEl.host}
                  onclick={({ target }) => {
                    if (!target.id.includes("heart")) {
                      choseSublistEl(listEl.host, service.serviceName);
                      lastChosenHost.set(listEl.host);
                      lastChosenService.set(service.serviceName);

                      if (!target.className.includes("log")) {
                        activeMenuOption.set("home");
                      }

                      initialVisitcounter = 1;
                    }
                  }}
                >
                  <div class="hostRow {customListElClass}">
                    <p
                      class={service.isDisabled ? "disabled" : ""}
                      title={service.serviceName}
                    >
                      {service.serviceName}
                    </p>
                    {#if listElementButton}
                      <div class="buttonBox flex">
                        <div
                          class="listElementButton"
                          onclick={() => {
                            navigate(
                              `${changeKey}/servicesettings/${listEl.host.trim()}/${service.serviceName.trim()}`,
                              { replace: true }
                            );

                            chosenElSettings = `${listEl.host.trim()}-${service.serviceName.trim()}`;
                          }}
                        >
                          <i class="log log-Wheel"></i>
                        </div>
                        <div
                          id={`heartButtonCont-${i}`}
                          class="listElementButton"
                          onclick={() => {
                            favoriteToggle(listEl.host, service.serviceName);
                          }}
                        >
                          <i
                            id={`heartButton-${i}`}
                            class="log {service.isFavorite
                              ? 'log-Heart'
                              : 'log-EmptyHeart'}"
></i>
                        </div>
                      </div>
                    {/if}
                  </div>
                  <div
                    class={`highlightedOverlay ${
                      `${activeElementName}` ===
                      `${listEl.host.trim()}-${service.serviceName.trim()}`
                        ? "active"
                        : ``
                    }`}
></div>
                </li>
              {/if}{/each}
          </ul>
          <div
            class="flex flex-start stopedServicesBox clickable inactiveServices {listEl.services.find(
              (e) => {
                return e.isDisabled;
              }
            )
              ? ''
              : 'visuallyHidden'}"
            onclick={() => {
              toggleArchivedVisible(index);
              initialVisitcounter = 1;
            }}
          >
            <i class="log log-Archive"></i>
            <p class="stopedServices">stoped services</p>
            <i
              class="log log-Pointer {!openStopedServIndexes.includes(index)
                ? 'rotated'
                : ''}"
></i>
          </div>

          <ul
            class="activeServices inactiveServices {!openStopedServIndexes.includes(
              index
            )
              ? 'visuallyHidden'
              : ''}"
          >
            {#each listEl.services as service, i}
              {#if service.isDisabled}<li
                  class="serviceListItem  "
                  id={listEl.host}
                  onclick={({ target }) => {
                    if (!target.id.includes("heart")) {
                      choseSublistEl(listEl.host, service.serviceName);
                      lastChosenHost.set(listEl.host);
                      lastChosenService.set(service.serviceName);
                      if (!target.className.includes("log")) {
                        activeMenuOption.set("home");
                      }

                      initialVisitcounter = 1;
                    }
                  }}
                >
                  <div class="hostRow {customListElClass}">
                    <p
                      class={service.isDisabled ? "disabled" : ""}
                      title={service.serviceName}
                    >
                      {service.serviceName}
                    </p>
                    {#if listElementButton}
                      <div class="buttonBox flex">
                        <div
                          class="listElementButton"
                          onclick={() => {
                            navigate(
                              `${changeKey}/servicesettings/${listEl.host.trim()}/${service.serviceName.trim()}`,
                              { replace: true }
                            );

                            chosenElSettings = `${listEl.host.trim()}-${service.serviceName.trim()}`;
                          }}
                        >
                          <i class="log log-Wheel"></i>
                        </div>
                        <div
                          id={`heartButtonContDissabled-${i}`}
                          class="listElementButton"
                          onclick={() => {
                            favoriteToggle(listEl.host, service.serviceName);
                          }}
                        >
                          <i
                            id={`heartButtonDissabled-${i}`}
                            class="log {service.isFavorite
                              ? 'log-Heart'
                              : 'log-EmptyHeart'}"
></i>
                        </div>
                      </div>
                    {/if}
                  </div>
                  <div
                    class={`highlightedOverlay ${
                      `${activeElementName}` ===
                      `${listEl.host.trim()}-${service.serviceName.trim()}`
                        ? "active"
                        : ``
                    }`}
></div>
                </li>{/if}{/each}
          </ul>
        </div>
      </li>
    {/each}
  </ul>
</div>
