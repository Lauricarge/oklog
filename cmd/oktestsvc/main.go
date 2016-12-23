package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	var (
		id   = flag.String("id", "foo", "ID for this instance")
		rate = flag.Int("rate", 5, "records per second")
	)
	flag.Parse()

	var (
		nBytes   uint64
		nRecords uint64
	)
	printRate := func(d time.Duration, printEvery int) {
		var prevBytes, prevRecords, iterationCount uint64
		for range time.Tick(d) {
			currBytes := atomic.LoadUint64(&nBytes)
			currRecords := atomic.LoadUint64(&nRecords)
			bytesPerSec := (float64(currBytes) - float64(prevBytes)) / d.Seconds()
			recordsPerSec := (float64(currRecords) - float64(prevRecords)) / d.Seconds()

			prevBytes = currBytes
			prevRecords = currRecords

			iterationCount++
			if iterationCount%uint64(printEvery) == 0 {
				fmt.Fprintf(os.Stderr, "%2ds average: %.2f bytes/sec, %.2f records/sec\n", int(d.Seconds()), bytesPerSec, recordsPerSec)
			}

		}
	}

	go printRate(1*time.Second, 10)
	go printRate(10*time.Second, 1)
	fmt.Fprintf(os.Stderr, "%s starting, %d records per second\n", *id, *rate)

	rand.Seed(time.Now().UnixNano())
	var count uint64
	hz := float64(time.Second) / float64(*rate)
	for range time.Tick(time.Duration(hz)) {
		count++
		if n, err := fmt.Fprintf(os.Stdout,
			"%s %s %d %s\n",
			time.Now().Format(time.RFC3339),
			*id,
			count,
			records[rand.Intn(len(records))],
		); err != nil {
			fmt.Fprintf(os.Stderr, "%d: %v\n", count, err)
		} else {
			atomic.AddUint64(&nBytes, uint64(n))
			atomic.AddUint64(&nRecords, 1)
		}
	}
}

var records = []string{
	"Once upon a midnight dreary, while I pondered, weak and weary,",
	"Over many a quaint and curious volume of forgotten lore—",
	"While I nodded, nearly napping, suddenly there came a tapping,",
	"As of some one gently rapping, rapping at my chamber door.",
	"“’Tis some visitor,” I muttered, “tapping at my chamber door—",
	"Only this and nothing more.”",
	"Ah, distinctly I remember it was in the bleak December;",
	"And each separate dying ember wrought its ghost upon the floor.",
	"Eagerly I wished the morrow;—vainly I had sought to borrow",
	"From my books surcease of sorrow—sorrow for the lost Lenore—",
	"For the rare and radiant maiden whom the angels name Lenore—",
	"Nameless here for evermore.",
	"And the silken, sad, uncertain rustling of each purple curtain",
	"Thrilled me—filled me with fantastic terrors never felt before;",
	"So that now, to still the beating of my heart, I stood repeating",
	"“’Tis some visitor entreating entrance at my chamber door—",
	"Some late visitor entreating entrance at my chamber door;—",
	"This it is and nothing more.”",
	"Presently my soul grew stronger; hesitating then no longer,",
	"“Sir,” said I, “or Madam, truly your forgiveness I implore;",
	"But the fact is I was napping, and so gently you came rapping,",
	"And so faintly you came tapping, tapping at my chamber door,",
	"That I scarce was sure I heard you”—here I opened wide the door;—",
	"Darkness there and nothing more.",
	"Deep into that darkness peering, long I stood there wondering, fearing,",
	"Doubting, dreaming dreams no mortal ever dared to dream before;",
	"But the silence was unbroken, and the stillness gave no token,",
	"And the only word there spoken was the whispered word, “Lenore?”",
	"This I whispered, and an echo murmured back the word, “Lenore!”—",
	"Merely this and nothing more.",
	"Back into the chamber turning, all my soul within me burning,",
	"Soon again I heard a tapping somewhat louder than before.",
	"“Surely,” said I, “surely that is something at my window lattice;",
	"Let me see, then, what thereat is, and this mystery explore—",
	"Let my heart be still a moment and this mystery explore;—",
	"’Tis the wind and nothing more!”",
	"Open here I flung the shutter, when, with many a flirt and flutter,",
	"In there stepped a stately Raven of the saintly days of yore;",
	"Not the least obeisance made he; not a minute stopped or stayed he;",
	"But, with mien of lord or lady, perched above my chamber door—",
	"Perched upon a bust of Pallas just above my chamber door—",
	"Perched, and sat, and nothing more.",
	"Then this ebony bird beguiling my sad fancy into smiling,",
	"By the grave and stern decorum of the countenance it wore,",
	"“Though thy crest be shorn and shaven, thou,” I said, “art sure no craven,",
	"Ghastly grim and ancient Raven wandering from the Nightly shore—",
	"Tell me what thy lordly name is on the Night’s Plutonian shore!”",
	"Quoth the Raven “Nevermore.”",
	"Much I marvelled this ungainly fowl to hear discourse so plainly,",
	"Though its answer little meaning—little relevancy bore;",
	"For we cannot help agreeing that no living human being",
	"Ever yet was blessed with seeing bird above his chamber door—",
	"Bird or beast upon the sculptured bust above his chamber door,",
	"With such name as “Nevermore.”",
	"But the Raven, sitting lonely on the placid bust, spoke only",
	"That one word, as if his soul in that one word he did outpour.",
	"Nothing farther then he uttered—not a feather then he fluttered—",
	"Till I scarcely more than muttered “Other friends have flown before—",
	"On the morrow he will leave me, as my Hopes have flown before.”",
	"Then the bird said “Nevermore.”",
	"Startled at the stillness broken by reply so aptly spoken,",
	"“Doubtless,” said I, “what it utters is its only stock and store",
	"Caught from some unhappy master whom unmerciful Disaster",
	"Followed fast and followed faster till his songs one burden bore—",
	"Till the dirges of his Hope that melancholy burden bore",
	"Of ‘Never—nevermore’.”",
	"But the Raven still beguiling all my fancy into smiling,",
	"Straight I wheeled a cushioned seat in front of bird, and bust and door;",
	"Then, upon the velvet sinking, I betook myself to linking",
	"Fancy unto fancy, thinking what this ominous bird of yore—",
	"What this grim, ungainly, ghastly, gaunt, and ominous bird of yore",
	"Meant in croaking “Nevermore.”",
	"This I sat engaged in guessing, but no syllable expressing",
	"To the fowl whose fiery eyes now burned into my bosom’s core;",
	"This and more I sat divining, with my head at ease reclining",
	"On the cushion’s velvet lining that the lamp-light gloated o’er,",
	"But whose velvet-violet lining with the lamp-light gloating o’er,",
	"She shall press, ah, nevermore!",
	"Then, methought, the air grew denser, perfumed from an unseen censer",
	"Swung by Seraphim whose foot-falls tinkled on the tufted floor.",
	"“Wretch,” I cried, “thy God hath lent thee—by these angels he hath sent thee",
	"Respite—respite and nepenthe from thy memories of Lenore;",
	"Quaff, oh quaff this kind nepenthe and forget this lost Lenore!”",
	"Quoth the Raven “Nevermore.”",
	"“Prophet!” said I, “thing of evil!—prophet still, if bird or devil!—",
	"Whether Tempter sent, or whether tempest tossed thee here ashore,",
	"Desolate yet all undaunted, on this desert land enchanted—",
	"On this home by Horror haunted—tell me truly, I implore—",
	"Is there—is there balm in Gilead?—tell me—tell me, I implore!”",
	"Quoth the Raven “Nevermore.”",
	"“Prophet!” said I, “thing of evil!—prophet still, if bird or devil!",
	"By that Heaven that bends above us—by that God we both adore—",
	"Tell this soul with sorrow laden if, within the distant Aidenn,",
	"It shall clasp a sainted maiden whom the angels name Lenore—",
	"Clasp a rare and radiant maiden whom the angels name Lenore.”",
	"Quoth the Raven “Nevermore.”",
	"“Be that word our sign of parting, bird or fiend!” I shrieked, upstarting—",
	"“Get thee back into the tempest and the Night’s Plutonian shore!",
	"Leave no black plume as a token of that lie thy soul hath spoken!",
	"Leave my loneliness unbroken!—quit the bust above my door!",
	"Take thy beak from out my heart, and take thy form from off my door!”",
	"Quoth the Raven “Nevermore.”",
	"And the Raven, never flitting, still is sitting, still is sitting",
	"On the pallid bust of Pallas just above my chamber door;",
	"And his eyes have all the seeming of a demon’s that is dreaming,",
	"And the lamp-light o’er him streaming throws his shadow on the floor;",
	"And my soul from out that shadow that lies floating on the floor",
	"Shall be lifted—nevermore!",

	"Turning and turning in the widening gyre",
	"The falcon cannot hear the falconer;",
	"Things fall apart; the centre cannot hold;",
	"Mere anarchy is loosed upon the world,",
	"The blood-dimmed tide is loosed, and everywhere",
	"The ceremony of innocence is drowned;",
	"The best lack all conviction, while the worst",
	"Are full of passionate intensity.",
	"Surely some revelation is at hand;",
	"Surely the Second Coming is at hand.",
	"The Second Coming! Hardly are those words out",
	"When a vast image out of Spiritus Mundi",
	"Troubles my sight: somewhere in sands of the desert",
	"A shape with lion body and the head of a man,",
	"A gaze blank and pitiless as the sun,",
	"Is moving its slow thighs, while all about it",
	"Reel shadows of the indignant desert birds.",
	"The darkness drops again; but now I know",
	"That twenty centuries of stony sleep",
	"Were vexed to nightmare by a rocking cradle,",
	"And what rough beast, its hour come round at last,",
	"Slouches towards Bethlehem to be born?",

	"I have been looking for you, my child,",
	"Since the time when rivers and mountains still lay in obscurity.",
	"I was looking for you when you were still in a deep sleep",
	"Although the conch had many times echoed in the ten directions.",
	"Without leaving our ancient mountain I looked at distant lands",
	"And recognized your steps on so many different paths.",
	"Where are you going, my child?",
	"There have been times when the mist has come",
	"And enveloped the remote village but you are still",
	"Wandering in far away lands.",
	"I have called your name with each breath,",
	"Confident that even though you have lost your",
	"Way over there you will finally find a way back to me.",
	"Sometimes I manifest myself right on the path",
	"You are treading but you still look at me as if I were a stranger",
	"You cannot see the connection between us in our",
	"Former lives you cannot remember the old vow you made.",
	"You have not recognized me",
	"Because your mind is caught up in images concerning a distant future.",
	"In former lifetimes you have often taken my hand",
	"and we have enjoyed walking together.",
	"We have sat together for a longtime at the foot of old pine trees.",
	"We have stood side by side in silence for hours",
	"Listening to the sound of the wind softly calling us",
	"And looking up at the while clouds floating by.",
	"You have picked up and given to me the firstred autumn leaf",
	"And I have taken you through forests deep in snow.",
	"But wherever we go we always return to our",
	"Ancient mountain to be near to the moon and stars",
	"To invite the big bell every morning to sound,",
	"And help living beings to wake up.",
	"We have sat quietly on the An Tu mountain’ with the",
	"Great Bamboo Forest Master",
	"Alongside the frangipani trees in blossom.",
	"We have taken boats out to sea to rescue the boat people as they drift.",
	"We have helped Master Van Hanh design the Thang",
	"Long capital we have built together a thatched hermitage,",
	"And stretched out the net to rescue the nun Trac Tuyen When!",
	"The sound of The rising tide was deafening",
	"On the banks of The Tien Duong river.",
	"Together we have opened the way and stepped",
	"Into the immense space outside of space.",
	"After many years of working to tear asunder the net of time.",
	"We have saved up the light of shooting stars",
	"And made a torch helping those who want to go home",
	"After decades of wandering in distant places.",
	"But still there have been times when the",
	"Seeds of a vagabond in you have come back to life",
	"you have left your teacher, your brothers and sisters",
	"Alone you go…",
	"I look at you with compassion",
	"Although I know that this is not a true separation",
	"(Because I am already in each cell of your body)",
	"And that you may need once more to play the prodigal son.",
	"That is why I promise I shall be there for you",
	"Any time you are in danger.",
	"Sometimes you have lain unconscious on the hot sands of frontier deserts.",
	"I have manifested myself as a cloud to bring you cool shade.",
	"Late at night the cloud became the dew",
	"And the compassionate nectar falls drop by drop for you to drink.",
	"Sometimes you sit in a deep abyss of darkness",
	"Completely alienated from you true home.",
	"I have manifested Myself as a long ladder and",
	"Lightly thrown myself down",
	"So that you can climb up to the area where there is light",
	"To discover again the blue of the sky and the",
	"Sounds of the brook and the birds.",
	"Sometimes I recognised you in Birmingham,",
	"In the Do Linh district or New England.",
	"I have sometimes met you in Hang Chau, Xiamen, or Shanghai",
	"I have sometimes found you in St. Petersburg or East Berlin.",
	"Sometimes, though only five years old, I have",
	"Seen you and recognized you.",
	"Because of the seed of bodhchita, you carry in your tender heart.",
	"Wherever I have seen you, I have always raised",
	"My hand and made a signal to you,",
	"Whether it be in the delta of the North, Saigon or the Thuan An Seaport.",
	"Sometimes you were the golden full moon hanging",
	"Over the summit of The Kim Son Mountain,",
	"Or the little bird flying over the Dai Laoforest during a winter night.",
	"Often I have seen you",
	"But you have not seen me,",
	"Though while walking in the evening mist your clothes have been soaked.",
	"But finally you have always come home.",
	"You have come home and sat at my feet on our ancient mountain",
	"Listening to the birds calling and the monkeys",
	"Screeching and the morning chanting echoing from the Buddha Hall.",
	"You have come back to me determined not to be a vagabond any longer.",
	"This morning the birds of the mountain joyfully welcome the bright sun.",
	"Do you know, my child, that the white clouds",
	"Are still floating in the vault of the sky?",
	"Where are you now?",
	"The ancient mountain is still there in this",
	"Place of the present moment.",
	"Although the white-crested wave still wants to",
	"Go in the other direction,",
	"Look again, you will see me in you and in every leaf and flower bud.",
	"If you call my name, you will see me right away.",
	"Where are you going?",
	"The old frangipani tree offers its fragrant flowers this morning.",
	"You and I have never really been apart. Spring has come.",
	"The pines have put out new shining green needles",
	"And on the edge of the forest, the wild Plum",
	"Trees have burst into flower.",

	"Before there was a trace of this world of men,",
	"I carried the memory of a lock of your hair,",
	"A stray end gathered within me, though unknown.",
	"Inside that invisible realm,",
	"Your face like the sun longed to be seen,",
	"Until each separate object was finally flung into light.",
	"From the moment of Time’s first-drawn breath,",
	"Love resides in us,",
	"A treasure locked into the heart’s hidden vault;",
	"Before the first seed broke open the rose bed of Being,",
	"An inner lark soared through your meadows,",
	"Heading toward Home.",
	"What can I do but thank you, one hundred times?",
	"Your face illumines the shrine of Hayati’s eyes,",
	"Constantly present and lovely.",

	"My house is buried in the deepest recess of the forest.",
	"Every year, ivy vines grow longer than the year before.",
	"Undisturbed by the affairs of the world I live at ease,",
	"Woodmen’s singing rarely reaching me through the trees.",
	"While the sun stays in the sky, I mend my torn clothes,",
	"And facing the moon, I read holy texts aloud to myself.",
	"Let me drop a word of advice for believers of my faith.",
	"To enjoy life’s immensity, you do not need many things.",
}