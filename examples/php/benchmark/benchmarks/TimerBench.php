<?php

use Example\Timer;

class TimerBench
{
    /**
     * @Revs(1000)
     * @Iterations(5)
     */
    public function benchConsume(): void
    {
       $consumer = new Timer();
       $consumer->consume();
    }
}
