<script>
  import { getContext } from 'svelte';
  import { derived } from 'svelte/store';
  export let path = '';
  export let component = null;

  const router = getContext('svelte-routing') || {};
  const location = (router.location && router.location) || { subscribe: (fn)=>{ fn({pathname: window.location.pathname}); return ()=>{} } };
  const basepath = router.basepath || '';

  function matchPath(pattern, pathname) {
    if (!pattern) return { params: {} };
    // remove basepath prefix
    if (basepath && pathname.startsWith(basepath)) {
      pathname = pathname.slice(basepath.length) || '/';
    }
    const p = pattern.replace(/(^\/|\/$)/g, '')
    const pa = p === '' ? [] : p.split('/')
    const na = pathname.replace(/(^\/|\/$)/g, '') === '' ? [] : pathname.replace(/(^\/|\/$)/g, '').split('/')
    if (pa.length !== na.length) return null
    const params = {}
    for (let i = 0; i < pa.length; i++) {
      if (pa[i].startsWith(':')) {
        params[pa[i].slice(1)] = decodeURIComponent(na[i])
      } else if (pa[i] !== na[i]) {
        return null
      }
    }
    return { params }
  }

  const match = derived(location, $loc => {
    return matchPath(path, $loc.pathname);
  });
</script>

{#if $match}
  {#if component}
    <svelte:component this={component} {...$match.params} />
  {:else}
    <slot />
  {/if}
{/if}
