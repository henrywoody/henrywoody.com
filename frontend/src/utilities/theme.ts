export function isDarkModeOn() {
	return window.matchMedia("(prefers-color-scheme: dark)").matches;
}
