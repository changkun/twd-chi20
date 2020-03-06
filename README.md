# The Walking Data - CHI2020 Late Breaking Work

_Kai Holländer, Luca Schellenberg, Changkun Ou, Andreas Butz_

Automated cars will need to observe pedestrians and react adequately to their behavior when driving in urban areas. Judging pedestrian behavior, however, is hard. When approaching it by machine learning methods, large amounts of training data is needed, which is costly and difficult to obtain, especially for critical situations. In order to provide such data, we have developed an online game inspired by Frogger, in which players have to cross streets. Accidents and critical situations are a natural part of the data produced in such a way without anybody getting hurt in reality. We present the design of our game and an analysis of the resulting data and its match to real world behavior observed in previous work. We found that behavior patterns in real and virtual environments correlated and argue that game data could be used to train machine learning algorithms for predicting real pedestrian's walking trajectories when crossing a road. This approach could be used in future automated vehicles to increase pedestrian safety.

## About this repo

This repository contains source code regarding a simple email subscription service.
The server side validates a submitted email, stores it in the BoltDB. Then send an email with given template.

## Usage

The server is written in Go. After the installation of Go, to run the server, 
all you need is:

```
make
make s
```

## Citation

If you find the project useful, please consider citing our paper

```
@inproceedings{chi2020-lbw-twd,
 author = {Holl\"{a}nder, Kai and Schelleneberg, Luca and Ou, Changkun and Butz, Andreas},
 title = {All Fun and Games: Obtaining Critical Pedestrian Behavior Data from an Online Simulation},
 booktitle = {Extended Abstracts, CHI Conference on Human Factors in Computing Systems},
 series = {CHI'20 Adjunct Proceedings},
 year = {2020},
 isbn = {978-1-4503-6819-3/20/04},
 location = {Honolulu, HI, USA},
 doi = {10.1145/3334480.3382797},
 publisher = {ACM},
 address = {New York, NY, USA},
 keywords = {Games with a purpose; pedestrian safety; machine learning; automated vehicles; user behavior}
}
```

## Licnese

GPLv3 &copy; Changkun Ou