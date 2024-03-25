/**
 * Creates an observer on a given element
 * @param element the element the observer is observing
 * @param fetchCallback the function to run when the observer sees the element
 * @returns an empty promise
 */
export function CreateObserver(element: HTMLElement | null, fetchCallback: () => Promise<void>) {
    if (element) {
        const observer = new IntersectionObserver(async (entries) => {
            const first = entries[0];
            if (first.isIntersecting) {
                //console.debug("Observer:: Observer intersection")
                await fetchCallback();
            }
        });
        observer.observe(element)

        return observer
    }
    return null
}