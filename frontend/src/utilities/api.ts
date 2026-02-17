export type AsteroidScore = {
	name: string;
	score: number;
};

const API_BASE_URL = import.meta.env.PUBLIC_API_BASE_URL;

export async function getAsteroidsScores(scoreCount?: number): Promise<AsteroidScore[]> {
	let url = API_BASE_URL + "/asteroids-scores";
	if (scoreCount) {
		url += `?limit=${scoreCount}`;
	}
	const response = await fetch(url);
	return await response.json();
}

export async function postAsteroidsScore(name: string, score: number) {
	const response = await fetch(API_BASE_URL + "/asteroids-scores", {
		method: "POST",
		headers: { "Content-Type": "application/json" },
		body: JSON.stringify({ name: name.toUpperCase(), score }),
	});
	return await response.json();
}
