//A daily entry can have multiple workouts
interface Entry {
  date: string//date
  workouts: Workout[]
}

interface Workout {
  summary: string;//text body
  name: string;
  blocks: Exercise[];
}

interface Exercise {
  block: number//??
  name: string// example - Deadlift, Backsquat, Row, Bike, BallSlam
  time?: number;//min?
  weight?: number;
  sets?: number;
  reps?: number;
  rest?: number;//secs?
}

[
  {
  },
  {
    name: 'Deadlift',
    time: null,
    weight: 225,
    set: 2,
    reps: 5,
    rest: 30,
  },
  {
    name: 'Deadlift',
    time: null,
    weight: 315,
    set: 2,
    reps: 5,
    rest: 120,
  }, 
  {
    name: 'Deadlift',
    time: null,
    weight: 345,
    set: 3,
    reps: 5,
    rest: 120,
  }
]

//Summary
/*
Block 1
Deadlift

2 sets of 10 at 135lbs with 30 sec of rest
2 sets of 5 at 225lbs with 30 sec of rest
2 sets of 5 at 315lbs with 2min of rest
3 sets of 5 at 345lbs with 2min of rest
*/

/*
  Block 2

 */