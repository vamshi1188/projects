import { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ArrowLeft, Check } from "lucide-react";
import type { ServiceType } from "./HomePage";

import haircutFade from "@assets/generated_images/Classic_fade_haircut_789ad47e.png";
import haircutPompadour from "@assets/generated_images/Pompadour_hairstyle_3a56d87f.png";
import haircutBuzz from "@assets/generated_images/Buzz_cut_style_948bb90d.png";
import haircutCrop from "@assets/generated_images/Textured_crop_haircut_ce4770a9.png";

import beardFull from "@assets/generated_images/Full_beard_style_b3eafe16.png";
import beardGoatee from "@assets/generated_images/Goatee_beard_style_e3d30f8d.png";
import beardStubble from "@assets/generated_images/Stubble_beard_style_93cd2fb7.png";
import beardMustache from "@assets/generated_images/Classic_mustache_style_bc06d84f.png";

import colorBalayage from "@assets/generated_images/Brown_balayage_color_94eb2a29.png";
import colorBlonde from "@assets/generated_images/Platinum_blonde_color_88795eb9.png";
import colorGray from "@assets/generated_images/Ash_gray_highlights_450066bc.png";
import colorMatched from "@assets/generated_images/Matched_hair_beard_color_1541c7a5.png";

import massageHead from "@assets/generated_images/Head_massage_service_22178beb.png";
import massageShoulder from "@assets/generated_images/Shoulder_massage_service_f7675610.png";
import massageStone from "@assets/generated_images/Hot_stone_massage_1d8ad678.png";

import facewashCleansing from "@assets/generated_images/Facial_cleansing_service_55369e12.png";
import facewashDeep from "@assets/generated_images/Deep_cleansing_facial_4d1da52f.png";
import facewashCharcoal from "@assets/generated_images/Charcoal_face_mask_70bbecd3.png";
import facewashCucumber from "@assets/generated_images/Refreshing_cucumber_facial_670ad394.png";

export interface Style {
  id: string;
  name: string;
  description: string;
  price: number;
  image: string;
}

const serviceStyles: Record<ServiceType, Style[]> = {
  haircut: [
    { id: "fade", name: "Classic Fade", description: "Timeless short sides with textured top", price: 500, image: haircutFade },
    { id: "pompadour", name: "Pompadour", description: "Sleek and voluminous swept back style", price: 700, image: haircutPompadour },
    { id: "buzz", name: "Buzz Cut", description: "Clean and minimalist short cut", price: 300, image: haircutBuzz },
    { id: "crop", name: "Textured Crop", description: "Modern messy fringe style", price: 600, image: haircutCrop },
  ],
  beard: [
    { id: "full", name: "Full Beard", description: "Complete beard shaping and styling", price: 300, image: beardFull },
    { id: "goatee", name: "Goatee", description: "Classic chin and mustache styling", price: 250, image: beardGoatee },
    { id: "stubble", name: "Stubble", description: "Short even facial hair trim", price: 200, image: beardStubble },
    { id: "mustache", name: "Mustache", description: "Upper lip grooming", price: 150, image: beardMustache },
  ],
  color: [
    { id: "balayage", name: "Brown Balayage", description: "Natural-looking brown highlights", price: 2500, image: colorBalayage },
    { id: "blonde", name: "Platinum Blonde", description: "Full platinum blonde color", price: 3000, image: colorBlonde },
    { id: "gray", name: "Ash Gray", description: "Contemporary silver highlights", price: 2800, image: colorGray },
    { id: "matched", name: "Hair & Beard Match", description: "Coordinated hair and beard coloring", price: 3500, image: colorMatched },
  ],
  massage: [
    { id: "head", name: "Head Massage", description: "Relaxing scalp massage therapy", price: 400, image: massageHead },
    { id: "shoulder", name: "Shoulder & Neck", description: "Upper body tension relief", price: 600, image: massageShoulder },
    { id: "stone", name: "Hot Stone", description: "Luxurious heated stone therapy", price: 1000, image: massageStone },
  ],
  facewash: [
    { id: "cleansing", name: "Facial Cleansing", description: "Gentle foam face wash treatment", price: 350, image: facewashCleansing },
    { id: "deep", name: "Deep Cleanse", description: "Pore cleansing exfoliation", price: 500, image: facewashDeep },
    { id: "charcoal", name: "Charcoal Mask", description: "Detoxifying charcoal treatment", price: 600, image: facewashCharcoal },
    { id: "cucumber", name: "Cucumber Refresh", description: "Hydrating cucumber facial", price: 450, image: facewashCucumber },
  ],
};

const serviceNames: Record<ServiceType, string> = {
  haircut: "Haircut",
  beard: "Beard",
  color: "Hair & Beard Color",
  massage: "Massage",
  facewash: "Face Wash",
};

interface ServiceSelectionProps {
  serviceType: ServiceType;
  onBack: () => void;
  onSkip: () => void;
  onContinue: (selectedStyles: Style[]) => void;
}

export default function ServiceSelection({ serviceType, onBack, onSkip, onContinue }: ServiceSelectionProps) {
  const [selectedStyles, setSelectedStyles] = useState<Set<string>>(new Set());
  const styles = serviceStyles[serviceType];

  const toggleStyle = (styleId: string) => {
    const newSelected = new Set(selectedStyles);
    if (newSelected.has(styleId)) {
      newSelected.delete(styleId);
    } else {
      newSelected.add(styleId);
    }
    setSelectedStyles(newSelected);
  };

  const handleContinue = () => {
    const selected = styles.filter(s => selectedStyles.has(s.id));
    onContinue(selected);
  };

  return (
    <div className="min-h-screen flex flex-col bg-background">
      <header className="sticky top-0 z-10 bg-background/95 backdrop-blur border-b">
        <div className="flex items-center justify-between p-4">
          <div className="flex items-center gap-2">
            <Button variant="ghost" size="icon" onClick={onBack} data-testid="button-back">
              <ArrowLeft className="h-5 w-5" />
            </Button>
            <h2 className="text-lg font-semibold">{serviceNames[serviceType]}</h2>
          </div>
          <Button variant="ghost" onClick={onSkip} data-testid="button-skip">
            Skip
          </Button>
        </div>
      </header>

      <div className="flex-1 p-6 space-y-4">
        {styles.map((style) => {
          const isSelected = selectedStyles.has(style.id);
          return (
            <Card
              key={style.id}
              className={`overflow-hidden hover-elevate active-elevate-2 cursor-pointer ${
                isSelected ? "ring-2 ring-primary" : ""
              }`}
              onClick={() => toggleStyle(style.id)}
              data-testid={`card-style-${style.id}`}
            >
              <div className="aspect-[4/3] relative">
                <img src={style.image} alt={style.name} className="w-full h-full object-cover" />
                {isSelected && (
                  <div className="absolute top-2 right-2 h-8 w-8 rounded-full bg-primary flex items-center justify-center">
                    <Check className="h-5 w-5 text-primary-foreground" />
                  </div>
                )}
              </div>
              <div className="p-4 space-y-2">
                <div className="flex items-start justify-between gap-2">
                  <div className="flex-1">
                    <h3 className="font-semibold text-lg">{style.name}</h3>
                    <p className="text-sm text-muted-foreground">{style.description}</p>
                  </div>
                  <Badge variant="secondary" className="flex-shrink-0">₹{style.price}</Badge>
                </div>
              </div>
            </Card>
          );
        })}
      </div>

      {selectedStyles.size > 0 && (
        <div className="sticky bottom-0 p-4 bg-background border-t">
          <Button className="w-full" onClick={handleContinue} data-testid="button-continue">
            Continue ({selectedStyles.size} selected)
          </Button>
        </div>
      )}
    </div>
  );
}
