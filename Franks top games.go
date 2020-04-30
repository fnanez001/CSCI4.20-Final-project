// Frank Nanez
// 4-25-20
// program to display my top 25 games played on steam and give information for user to learn more about the game.

package main

import "fmt"

type ginfo struct {
  Number  int
  Name, Hrs, Genre, Synopsis string

}

func game(i int) (int, string, int, string, string) {



  game1 = ginfo{
    Number: 1,
    Name: "Banished",
    Hrs: 349,
    Genre: "Simulation", "Stragey", "City Builder",
    Synopsis: "You and a group of exiled travelers who decide to restart their lives in a new land. They have only the clothes on their backs and a cart filled with supplies from their homeland. The townspeople of Banished are your primary resource. You must build, survive, and if you’re good enough, thrive.",
  }
  game2 = ginfo{
    Number: 2,
    Name:	"Dota 2",
    Hrs: 142,
    Genre: "MOBA, strategy",
    Synopsis: "Primary game mode is player vs player (PvP) where two teams of five battle against one another with waves of computer minions to help you push a “lane” guarded by other players, their minions, and towers.",
  }
  game3 = ginfo{
    Number: 3,
    Name: "Mini Ninja’s",
    Hrs: 94,
    Genre: "action, adventure, indie",
    Synopsis: "Mini Ninjas is a game that combines furious action with stealth and exploration for an experience that appeals to a wide audience across age groups and preferences. Set with a ninja oriental theme. Fight against the darkness that is attacking your home and the ones you know and love.",
  }
  game4 = ginfo{
    Number: 4,
    Name: "The Elder Scrolls V: Skyrim",
    Hrs: 79,
    Genre: "Open world, RPG, first person",
    Synopsis: "Dragons were thought to be extinct. You, a prisoner, survive and attack by a dragon and find out that you are a “dragonborn”. One who can speak the dragon language and use their shouts, a special ability/magic. You can create a character with tons of customization options and explore and play the game anyway you desire with the vast open world with plenty of choice and quests to take.",
  }
  game5 = ginfo{
    Number: 5,
    Name: "Sid Meier's Civilization V",
    Hrs: 62,
    Genre: "Turn-based strategy, 4x",
    Synopsis: "Play one of history’s greatest civilizations. Employ diplomatic, research, and military strategies to be the best civilization and win, choosing from many maps and map sizes for the set of your theatre.",
  }
  game6 = ginfo{
    Number: 6,
    Name: "Borderlands 2",
    Hrs: 55,
    Genre: "open world, RPG, first person",
    Synopsis: "Play as one of four vault hunters facing off against a massive world of creatures, psychos and the evil mastermind, Handsome Jack. Make new friends, arm them with a bazillion weapons to help you on your adventure.",
  }
  game7 = ginfo{
    Number: 7,
    Name: "Company of Heroes-Legacy edition",
    Hrs: 51,
    Genre: "real-time strategy",
    Synopsis: "Delivering a visceral WWII gaming experience, Company of Heroes redefines real time strategy gaming by bringing the sacrifice of heroic soldiers, war-ravaged environments, and dynamic battlefields to life. Beginning with the D-Day Invasion of Normandy, players lead squads of Allied soldiers into battle against the German war machine through some of the most pivotal battles of WWII. Through a rich single player campaign, players experience the cinematic intensity and bravery of ordinary soldiers thrust into extraordinary events.",
  }
  game8 = ginfo{
    Number: 8,
    Name: "PAYDAY 2",
    Hrs: 47,
    Genre: "First person shooter, strategy/heist",
    Synopsis: "Team up with up to four friends or other online players to take on heists and earn money, buy weapons, and customize your gear to help you be successful.",
  }
  game9 = ginfo{
    Number: 9,
    Name: "9.	The Forest",
    Hrs: 47,
    Genre: "Survival, horror, PVE",
    Synopsis: "Your plane crash lands on this beautiful island with a lush environment. You awake after the crash landing, only remembering your son calling out to you as he is dragged away, having blurry vision you don’t remember many details. You need to build some shelter and a base camp for your exploration. Soon after establishing some semblance of a shelter you see something off in the distance scoping out your new home. It’s a savage tribal inhabitant that you soon find out is a cannibal. You need to hurry and find your son and find a way off the island.",
  }
  game10 = ginfo{
    Number: 10,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game11 = ginfo{
    Number: 11,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game12 = ginfo{
    Number: 12,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game13 = ginfo{
    Number: 13,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game14 = ginfo{
    Number: 14,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game15 = ginfo{
    Number: 15,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game16 = ginfo{
    Number: 16,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game17 = ginfo{
    Number: 17,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game18 = ginfo{
    Number: 18,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game19 = ginfo{
    Number: 19,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game20 = ginfo{
    Number: 20,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game21 = ginfo{
    Number: 21,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game22 = ginfo{
    Number: 22,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game23 = ginfo{
    Number: 23,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game24 = ginfo{
    Number: 24,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  game25 = ginfo{
    Number: 25,
    Name:
    Hrs: 
    Genre:
    Synopsis:
  }
  return game(n) Number, Name, Hrs, Genre, Synopsis
}

func main (){
  
  x := 0

  for x != "Done" OR x != "done" OR x != "DONE"
    for i = 0, i <= 25, i++{
      x = game(g[x].Number,game(g[x].Name))
      fmt.Println(game(g[x].Number, game(g[x].Name)))
      x = x+1
    }

  


}