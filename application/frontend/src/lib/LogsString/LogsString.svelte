<script>
  import { tryToParseLogString } from "../../utils/functions";
  import { chosenStatus } from "../../Stores/stores";
  import fetchApi from "../../utils/fetch";

  import { store } from "../../Stores/stores.js";
  /**
   * @typedef {Object} Props
   * @property {string} [status]
   * @property {string} [time]
   * @property {string} [message]
   * @property {string} [width]
   * @property {boolean} [isHiglighted]
   * @property {any} [sharedLinkCallBack]
   * @property {any} [getLogsByTagOptions]
   */

  /** @type {Props} */
  let {
    status = "",
    time = "",
    message = "",
    width = "",
    isHiglighted = false,
    sharedLinkCallBack = () => {},
    getLogsByTagOptions = {}
  } = $props();

  let activeStatus = "";
  let parsedStr = $derived(tryToParseLogString(message));
</script>

<tr
  class="logsString {isHiglighted ? 'new' : ''} {message?.trim().length === 0
    ? 'emptyLogsString'
    : ''}"
  style="width: {width}px"
>
  <td
    onclick={async () => {
      if ($chosenStatus !== status) {
        chosenStatus.set(status);
      } else {
        chosenStatus.set("");
      }
    }}
    class="status {status ? status : 'hidden'} {status === $chosenStatus
      ? 'chosenStatus'
      : ''}"><p><span> ◉ </span>{status.toUpperCase()}</p></td
  >

  <td class="time row_group"
    ><p>{message?.trim()?.length > 0 ? time : ""}</p>
    <div>
      {#if message?.trim()?.length > 0}
        <div
          id={`thumb-shared-${time}`}
          class="shareLinkButtonThumb"
          onclick={() => {
            sharedLinkCallBack();
          }}
        >
          <i class="log log-ShareLink" id={`shared-${time}`}></i>
        </div>{/if}
    </div>
  </td>
  <td class="message">
    {#if !parsedStr}<p>
        {@html message}
      </p>{:else if $store.transformJson}<p>{parsedStr.startText}</p>
      <pre>{@html parsedStr.html}</pre>
      <p>{parsedStr.endText}</p>
    {:else}<p>
        {@html message}
      </p>{/if}
  </td>
</tr>
